package main

import (
	"fmt"
	"testing"

	"github.com/shadesfear/aoc-lib-go/datastructures"
	"github.com/shadesfear/aoc-lib-go/files"
)

func TestAVertical(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	grid, err := files.ParseInputToGrid(input)
	if err != nil {
		t.Error(err)
	}
	// target := []rune{'X', 'M', 'A', 'S'}
	target := []rune{'S', 'A', 'M', 'X'}

	ok := check(grid, datastructures.NewPoint(9, 6), datastructures.NewPoint(0, 1), target)

	fmt.Println(ok)
}

// func TestADiag(t *testing.T) {
// 	input := `MMMSXXMASM
// MSAMXMSMSA
// AMXSXMAAMM
// MSAMASMSMX
// XMASAMXAMM
// XXAMMXXAMA
// SMSMSASXSS
// SAXAMASAAA
// MAMMMXMMMM
// MXMXAXMASX`
// 	grid, err := files.ParseInputToGrid(input)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	target := []rune{'X', 'M', 'A', 'S'}
//
// 	ok := checkADiag(grid, datastructures.NewPoint(4, 0), target)
//
// 	fmt.Println(ok)
// }
//
// func TestAForward(t *testing.T) {
// 	input := `MMMSXXMASM
// MSAMXMSMSA
// AMXSXMAAMM
// MSAMASMSMX
// XMASAMXAMM
// XXAMMXXAMA
// SMSMSASXSS
// SAXAMASAAA
// MAMMMXMMMM
// MXMXAXMASX`
// 	grid, err := files.ParseInputToGrid(input)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	target := []rune{'X', 'M', 'A', 'S'}
//
// 	ok := checkAForward(grid, datastructures.NewPoint(5, 0), target)
//
// 	fmt.Println(ok)
// }

func TestPart1(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	grid, err := files.ParseInputToGrid(input)
	if err != nil {
		t.Error(err)
	}

	expected := 18

	result := solvePart1(grid)

	if expected != result {
		t.Errorf("Part1:  expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	grid, err := files.ParseInputToGrid(input)
	if err != nil {
		t.Error(err)
	}

	expected := 9

	result := solvePart2(grid)

	if expected != result {
		t.Errorf("Part2:  expected %d, got %d", expected, result)
	}
}
