package main

import (
	"testing"
)

func TestPart1Dummy1(t *testing.T) {
	input := `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`

	exp := 2
	res := solvePart1(input)
	if exp != res {
		t.Errorf("Part1, exp: %d, got %d", exp, res)
	}
}

func TestPart1Dummy2(t *testing.T) {
	input := `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`

	exp := 4
	res := solvePart1(input)
	if exp != res {
		t.Errorf("Part1, exp: %d, got %d", exp, res)
	}
}

func TestPart1(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	exp := 36
	res := solvePart1(input)
	if exp != res {
		t.Errorf("Part1, exp: %d, got %d", exp, res)
	}
}

func TestPart2(t *testing.T) {}
