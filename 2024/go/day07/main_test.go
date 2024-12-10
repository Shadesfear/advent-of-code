package main

import (
	"testing"

	"github.com/shadesfear/aoc-lib-go/files"
)

func TestPart1(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	lines := files.StringToLines(input)
	ex := 3749
	res := solvePart1(lines)
	if ex != res {
		t.Errorf("Got %d expected %d", res, ex)
	}
}

func TestPart2(t *testing.T) {}