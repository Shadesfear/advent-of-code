package main

import (
	"testing"

	"github.com/shadesfear/aoc-lib-go/files"
)

func TestPart1(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	lines := files.StringToLines(input)
	expected := 2
	result := solvePart1(lines)
	if result != expected {
		t.Errorf("Part1: got %d, expected %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	lines := files.StringToLines(input)
	expected := 4
	result := solvePart2(lines)
	if result != expected {
		t.Errorf("Part2: got %d, expected %d", result, expected)
	}
}

func TestPart3(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	lines := files.StringToLines(input)
	expected := 4
	result := solvePart2Clever(lines)
	if result != expected {
		t.Errorf("Part2: got %d, expected %d", result, expected)
	}
}
