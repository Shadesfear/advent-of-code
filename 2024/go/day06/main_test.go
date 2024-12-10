package main

import (
	// "log"
	// "log"
	"testing"

	"github.com/shadesfear/aoc-lib-go/files"
	// "github.com/shadesfear/aoc-lib-go/files"
	// "github.com/shadesfear/aoc-lib-go/files"
)

func TestPart1(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	exp := 41
	res := solvePart1(input)
	if exp != res {
		t.Errorf("part 1, Got: %d expected %d", res, exp)
	}
}

// func TestLoopDetector(t *testing.T) {
// 	input := `....#.....
// .........#
// ..........
// ..#.......
// .......#..
// ..........
// .#..^.....
// ........#.
// #.........
// ......#...`
//
// 	grid, err := files.ParseInputToGrid(input)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	exp := true
// 	start, dir := findStartPos(grid)
// 	res := loopDetector()
// 	if exp != res {
// 		t.Errorf("part 2, Got: %d expected %d", res, exp)
// 	}
// }

func TestPart2(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	exp := 6
	res := solvePart2(input)
	if exp != res {
		t.Errorf("part 2, Got: %d expected %d", res, exp)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines, _ := files.ReadInputFile("input.txt")

	solvePart1(lines)
	solvePart2(lines)
}
