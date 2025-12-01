package main

import (
	"testing"

	"github.com/shadesfear/aoc-lib-go/files"
)

func TestPart1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	lines := files.StringToLines(input)

	exp := 3
	res := solvePart1(lines)
	if res != exp {
		t.Errorf("solvePart1() = %d, exp %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	lines := files.StringToLines(input)
	exp := 6
	res := solvePart2(lines)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}
