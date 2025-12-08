package main

import (
	"testing"

	"github.com/shadesfear/aoc-lib-go/files"
)

func TestPart1(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	exp := 3
	res := solvePart1(input)
	if res != exp {
		t.Errorf("solvePart1() = %d, exp %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	exp := 14
	res := solvePart2(input)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}

func BenchmarkPark2(b *testing.B) {
	input, _ := files.ReadDayInput(5)
	db := Parse(input)

	for b.Loop() {
		SolvePart2Parsed(db)
	}
}
