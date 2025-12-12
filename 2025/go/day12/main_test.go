package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`

	exp := 2
	res := solvePart1(input)
	if res != exp {
		t.Errorf("solvePart1() = %d, exp %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	input := ""
	exp := 0
	res := solvePart2(input)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}
