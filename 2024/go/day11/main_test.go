package main

import (
	// "log"
	"testing"

	"github.com/shadesfear/aoc-lib-go/files"
)

func TestPart1(t *testing.T) {
	line := `125 17`
	exp := 55312
	input, err := files.ParseInputToInts(line)
	if err != nil {
		t.Errorf("%s", err)
	}
	res := solvePart1(input)
	if res != exp {
		t.Errorf("solvePart1() = %d, exp %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	line := `125 17`
	exp := 55312

	input, err := files.ParseInputToInts(line)
	if err != nil {
		t.Errorf("%s", err)
	}
	res := solvePart2(input)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}
