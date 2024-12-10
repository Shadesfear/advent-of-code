package main

import (
	"testing"
	// "github.com/shadesfear/aoc-lib-go/files"
)

func TestPart1(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := 161
	result := solvePart1(input)

	if result != expected {
		t.Errorf("Part1 - got %d, expected %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := 48
	result := solvePart2Better(input)

	if result != expected {
		t.Errorf("Part2 - got %d, expected %d", result, expected)
	}
}
