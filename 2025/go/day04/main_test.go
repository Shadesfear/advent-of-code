package main

import (
	"testing"

	"github.com/shadesfear/aoc-lib-go/files"
)

func TestPart1(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	grid, _ := files.ParseInputToGrid(input)
	exp := 13
	res := solvePart1(grid)
	if res != exp {
		t.Errorf("solvePart1() = %d, exp %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	grid, _ := files.ParseInputToGrid(input)
	exp := 43
	res := solvePart2(grid)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}
