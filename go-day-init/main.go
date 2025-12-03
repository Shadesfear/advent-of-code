package main

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

//go:embed templates/main.go.tmpl
var MAIN string

//go:embed templates/main_test.go.tmpl
var TEST string

func main() {
	app := &cli.App{
		Name:  "gdi",
		Usage: "Initialize Advent of Code day scaffolding",
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "Create main.go and main_test.go only",
				Action: cmdInit,
			},
			{
				Name:   "input",
				Usage:  "Download puzzle input, puzzle text, and examples",
				Action: cmdInput,
			},
		},
		Action:    cmdAll,
		Args:      true,
		ArgsUsage: "[year] [day]",
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cmdAll(c *cli.Context) error {
	year, day, dayDir, err := resolveContext(c)
	if err != nil {
		return err
	}

	if err := ensureDirectories(dayDir, year, day); err != nil {
		return err
	}

	if err := os.Chdir(dayDir); err != nil {
		return fmt.Errorf("could not change to day directory: %w", err)
	}

	puzzle, examples, puzzleErr := fetchPuzzle(year, day)
	initErr := initFiles(year, day, examples)
	inputErr := downloadInput(year, day)

	if puzzleErr != nil {
		fmt.Printf("Warning: could not fetch puzzle: %v\n", puzzleErr)
	} else {
		if err := os.WriteFile("puzzle.md", []byte(puzzle), 0o644); err != nil {
			fmt.Printf("Warning: could not write puzzle.md: %v\n", err)
		} else {
			fmt.Println("Created puzzle.md")
		}
	}
	if initErr != nil {
		fmt.Printf("init: %v\n", initErr)
	}
	if inputErr != nil {
		fmt.Printf("input: %v\n", inputErr)
	}

	if initErr != nil && inputErr != nil {
		return fmt.Errorf("both init and input failed")
	}
	return nil
}

func cmdInit(c *cli.Context) error {
	year, day, dayDir, err := resolveContext(c)
	if err != nil {
		return err
	}

	if err := ensureDirectories(dayDir, year, day); err != nil {
		return err
	}

	if err := os.Chdir(dayDir); err != nil {
		return fmt.Errorf("could not change to day directory: %w", err)
	}

	puzzle, examples, puzzleErr := fetchPuzzle(year, day)
	if puzzleErr != nil {
		fmt.Printf("Warning: could not fetch puzzle: %v\n", puzzleErr)
	} else {
		if err := os.WriteFile("puzzle.md", []byte(puzzle), 0o644); err != nil {
			fmt.Printf("Warning: could not write puzzle.md: %v\n", err)
		} else {
			fmt.Println("Created puzzle.md")
		}
	}

	return initFiles(year, day, examples)
}

func cmdInput(c *cli.Context) error {
	year, day, dayDir, err := resolveContext(c)
	if err != nil {
		return err
	}

	if err := ensureDirectories(dayDir, year, day); err != nil {
		return err
	}

	if err := os.Chdir(dayDir); err != nil {
		return fmt.Errorf("could not change to day directory: %w", err)
	}

	inputErr := downloadInput(year, day)

	puzzle, examples, puzzleErr := fetchPuzzle(year, day)
	if puzzleErr != nil {
		fmt.Printf("Warning: could not fetch puzzle: %v\n", puzzleErr)
	} else {
		if err := os.WriteFile("puzzle.md", []byte(puzzle), 0o644); err != nil {
			fmt.Printf("Warning: could not write puzzle.md: %v\n", err)
		} else {
			fmt.Println("Created puzzle.md")
		}

		if len(examples) > 0 {
			if err := writeExamplesFile(examples); err != nil {
				fmt.Printf("Warning: could not write examples.txt: %v\n", err)
			} else {
				fmt.Printf("Created examples.txt (%d examples)\n", len(examples))
			}
		}
	}

	return inputErr
}

func resolveContext(c *cli.Context) (year, day int, dayDir string, err error) {
	args := c.Args().Slice()

	switch len(args) {
	case 0:
		year, day, err = parseYearAndDayFromCwd()
		if err != nil {
			return 0, 0, "", err
		}
		dayDir, _ = os.Getwd()
	case 1:
		day, err = strconv.Atoi(args[0])
		if err != nil {
			return 0, 0, "", fmt.Errorf("invalid day: %s", args[0])
		}
		year = time.Now().Year()
		dayDir, err = buildDayDir(year, day)
		if err != nil {
			return 0, 0, "", err
		}
	case 2:
		year, err = strconv.Atoi(args[0])
		if err != nil {
			return 0, 0, "", fmt.Errorf("invalid year: %s", args[0])
		}
		day, err = strconv.Atoi(args[1])
		if err != nil {
			return 0, 0, "", fmt.Errorf("invalid day: %s", args[1])
		}
		dayDir, err = buildDayDir(year, day)
		if err != nil {
			return 0, 0, "", err
		}
	default:
		return 0, 0, "", fmt.Errorf("too many arguments")
	}

	return year, day, dayDir, nil
}

func buildDayDir(year, day int) (string, error) {
	root, err := findRepoRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, fmt.Sprintf("%d", year), "go", fmt.Sprintf("day%02d", day)), nil
}

func findRepoRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := cwd
	for {
		if hasGitDir(dir) && hasYearDir(dir) {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("could not find repo root (looking for .git + year folder)")
}

func hasGitDir(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, ".git"))
	return err == nil
}

func hasYearDir(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	yearPattern := regexp.MustCompile(`^20\d{2}$`)
	for _, e := range entries {
		if e.IsDir() && yearPattern.MatchString(e.Name()) {
			return true
		}
	}
	return false
}

func ensureDirectories(dayDir string, year, day int) error {
	if err := os.MkdirAll(dayDir, 0o755); err != nil {
		return fmt.Errorf("could not create day directory: %w", err)
	}

	root, err := findRepoRoot()
	if err != nil {
		return err
	}
	inputsDir := filepath.Join(root, fmt.Sprintf("%d", year), "inputs")
	if err := os.MkdirAll(inputsDir, 0o755); err != nil {
		return fmt.Errorf("could not create inputs directory: %w", err)
	}

	return nil
}

func parseYearAndDayFromCwd() (int, int, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return 0, 0, fmt.Errorf("could not get working directory: %w", err)
	}

	dayPattern := regexp.MustCompile(`day(\d+)`)
	yearPattern := regexp.MustCompile(`(20\d{2})`)

	dayMatch := dayPattern.FindStringSubmatch(filepath.Base(cwd))
	if dayMatch == nil {
		return 0, 0, fmt.Errorf("could not find day in path: %s", cwd)
	}
	day, _ := strconv.Atoi(dayMatch[1])

	yearMatch := yearPattern.FindStringSubmatch(cwd)
	if yearMatch == nil {
		return 0, 0, fmt.Errorf("could not find year in path: %s", cwd)
	}
	year, _ := strconv.Atoi(yearMatch[1])

	return year, day, nil
}

func initFiles(year, day int, examples []string) error {
	mainExists := fileExists("main.go")
	testExists := fileExists("main_test.go")

	if mainExists && testExists {
		return fmt.Errorf("main.go and main_test.go already exist")
	}

	mainContent := strings.Replace(MAIN, "day01.txt", fmt.Sprintf("day%02d.txt", day), 1)

	if !mainExists {
		if err := os.WriteFile("main.go", []byte(mainContent), 0o644); err != nil {
			return fmt.Errorf("could not write main.go: %w", err)
		}
		fmt.Printf("Created main.go (year=%d, day=%d)\n", year, day)
	}

	if !testExists {
		testContent := TEST
		if len(examples) > 0 {
			testContent = injectExample(TEST, examples[0])
		}
		if err := os.WriteFile("main_test.go", []byte(testContent), 0o644); err != nil {
			return fmt.Errorf("could not write main_test.go: %w", err)
		}
		fmt.Println("Created main_test.go")
	}

	if len(examples) > 0 {
		if err := writeExamplesFile(examples); err != nil {
			fmt.Printf("Warning: could not write examples.txt: %v\n", err)
		} else {
			fmt.Printf("Created examples.txt (%d examples)\n", len(examples))
		}
	}

	return nil
}

func injectExample(testTemplate, example string) string {
	lines := strings.Split(example, "\n")
	var quoted []string
	for _, line := range lines {
		quoted = append(quoted, fmt.Sprintf("\t\t%q,", line))
	}
	replacement := "input := []string{\n" + strings.Join(quoted, "\n") + "\n\t}"
	return strings.Replace(testTemplate, "input := []string{\n\t\t\"\",\n\t}", replacement, 1)
}

func writeExamplesFile(examples []string) error {
	var sb strings.Builder
	for i, ex := range examples {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprintf("--- Example %d ---\n", i+1))
		sb.WriteString(ex)
		sb.WriteString("\n")
	}
	return os.WriteFile("examples.txt", []byte(sb.String()), 0o644)
}

func fetchPuzzle(year, day int) (puzzle string, examples []string, err error) {
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return "", nil, fmt.Errorf("AOC_SESSION not set")
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", nil, err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", nil, fmt.Errorf("status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	html := string(body)
	puzzle = extractPuzzleText(html)
	examples = extractCodeBlocks(html)
	return puzzle, examples, nil
}

func extractPuzzleText(html string) string {
	articlePattern := regexp.MustCompile(`(?s)<article[^>]*>(.*?)</article>`)
	matches := articlePattern.FindAllStringSubmatch(html, -1)

	var parts []string
	for _, m := range matches {
		text := htmlToMarkdown(m[1])
		if len(text) > 0 {
			parts = append(parts, text)
		}
	}
	return strings.Join(parts, "\n\n---\n\n")
}

func htmlToMarkdown(html string) string {
	text := html

	text = regexp.MustCompile(`(?s)<h2[^>]*>(.*?)</h2>`).ReplaceAllString(text, "## $1\n\n")
	text = regexp.MustCompile(`(?s)<p>(.*?)</p>`).ReplaceAllString(text, "$1\n\n")
	text = regexp.MustCompile(`(?s)<pre><code>(.*?)</code></pre>`).ReplaceAllString(text, "```\n$1\n```\n\n")
	text = regexp.MustCompile(`(?s)<code>(.*?)</code>`).ReplaceAllString(text, "`$1`")
	text = regexp.MustCompile(`(?s)<em[^>]*>(.*?)</em>`).ReplaceAllString(text, "**$1**")
	text = regexp.MustCompile(`(?s)<li>(.*?)</li>`).ReplaceAllString(text, "- $1\n")
	text = regexp.MustCompile(`(?s)<ul>(.*?)</ul>`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`(?s)<a[^>]*href="([^"]*)"[^>]*>(.*?)</a>`).ReplaceAllString(text, "[$2]($1)")
	text = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(text, "")

	text = decodeHTMLEntities(text)
	text = regexp.MustCompile(`\n{3,}`).ReplaceAllString(text, "\n\n")
	text = strings.TrimSpace(text)

	return text
}

func extractCodeBlocks(html string) []string {
	pattern := regexp.MustCompile(`(?s)<pre><code>([^<]+)</code></pre>`)
	matches := pattern.FindAllStringSubmatch(html, -1)

	var blocks []string
	for _, m := range matches {
		block := strings.TrimSpace(m[1])
		block = decodeHTMLEntities(block)
		if len(block) > 0 {
			blocks = append(blocks, block)
		}
	}
	return blocks
}

func decodeHTMLEntities(s string) string {
	replacements := map[string]string{
		"&lt;":   "<",
		"&gt;":   ">",
		"&amp;":  "&",
		"&quot;": "\"",
		"&#39;":  "'",
	}
	for entity, char := range replacements {
		s = strings.ReplaceAll(s, entity, char)
	}
	return s
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func downloadInput(year, day int) error {
	inputPath := fmt.Sprintf("../../inputs/day%02d.txt", day)

	if fileExists(inputPath) {
		fmt.Printf("Input already exists: %s\n", inputPath)
		return nil
	}

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return fmt.Errorf("AOC_SESSION environment variable not set")
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not fetch input: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response: %w", err)
	}

	inputDir := filepath.Dir(inputPath)
	if err := os.MkdirAll(inputDir, 0o755); err != nil {
		return fmt.Errorf("could not create inputs directory: %w", err)
	}

	if err := os.WriteFile(inputPath, body, 0o644); err != nil {
		return fmt.Errorf("could not write input file: %w", err)
	}

	fmt.Printf("Downloaded input to %s\n", inputPath)
	return nil
}
