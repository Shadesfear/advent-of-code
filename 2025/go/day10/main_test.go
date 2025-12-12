package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
		"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
		"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
	}
	exp := 7
	res := solvePart1(input)
	if res != exp {
		t.Errorf("solvePart1() = %d, exp %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
		"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
		"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
	}
	exp := 33
	res := solvePart2(input)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}
