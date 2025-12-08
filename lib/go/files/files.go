package files

import (
	"bufio"
	"fmt"
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

func ParseLinesToGrid(input []string) ([][]rune, error) {

	lines := []string{}

	for _, line := range input {
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

func ParseInputToGrid(input string) ([][]rune, error) {
	liness := strings.Split(input, "\n")
	grid, err := ParseLinesToGrid(liness)
	return grid, err
}

func ReadDayInput(day int) (string, error) {
	path := fmt.Sprintf("../../inputs/day%02d.txt", day)
	return ReadInputFile(path)
}

func ReadDayInputLines(day int) ([]string, error) {
	path := fmt.Sprintf("../../inputs/day%02d.txt", day)
	return ReadInputLines(path)
}
