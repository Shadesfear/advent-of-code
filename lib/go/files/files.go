package files

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadInputFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadInputLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(data), "\n")
	non_empty := []string{}
	for _, line := range lines {
		if len(line) > 0 {
			non_empty = append(non_empty, line)
		}
	}

	return non_empty, nil
}

func StringToLines(input string) []string {
	return strings.Split(input, "\n")
}

func ParseInputToInts(input string) ([]int, error) {
	var nums []int
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nums, nil
}

func ParseInputToGrid(input string) ([][]rune, error) {
	liness := strings.Split(input, "\n")
	lines := []string{}

	for _, line := range liness {
		if len(line) < 1 {
			continue
		}
		lines = append(lines, line)
	}

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		if len(line) < 1 {
			continue
		}
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
		}
	}

	return grid, nil
}
