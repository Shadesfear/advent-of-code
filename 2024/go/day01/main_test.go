package main

import (
	"testing"

	"github.com/shadesfear/aoc-lib-go/files"
)

func TestPart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	lines := files.StringToLines(input)
	extected := 11
	result := solvePart1(lines)
	if result != extected {
		t.Errorf("solvePart1() = %d, expected %d", result, extected)
	}
}

func TestPart2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	lines := files.StringToLines(input)
	extected := 31
	result := solvePart2(lines)
	if result != extected {
		t.Errorf("solvePart1() = %d, expected %d", result, extected)
	}
}

func BenchmarkPart1(b *testing.B) {
	lines, err := files.ReadInputLines("input.txt")
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		solvePart1(lines)
	}
}
