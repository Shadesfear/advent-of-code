package main

import (
	"log"

	ds "github.com/shadesfear/aoc-lib-go/datastructures"
	"github.com/shadesfear/aoc-lib-go/files"
)

var grid [][]rune

func main() {
	lines, err := files.ReadInputFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}

	grid, err = files.ParseInputToGrid(lines)
	if err != nil {
		log.Fatal(err)
	}

	part1 := solvePart1(grid)
	log.Println(part1)

	part2 := solvePart2(grid)
	log.Println(part2)
}

func check(grid [][]rune, start, dir ds.Point, target []rune) bool {
	ROWS := len(grid)
	COLS := len(grid[0])

	if !start.InBounds(ROWS, COLS) {
		log.Println("Not in bounds")
		return false
	}

	lp := ds.NewPoint(start.X, start.Y)
	for i := 0; i < len(target); i++ {

		if !lp.InBounds(ROWS, COLS) {
			return false
		}
		if grid[lp.Y][lp.X] != target[i] {
			return false
		}

		lp = lp.Move(dir.X, dir.Y)

	}

	return true
}

func solvePart1(grid [][]rune) int {
	res := 0

	target := []rune{'X', 'M', 'A', 'S'}
	targetReverse := []rune{'S', 'A', 'M', 'X'}

	directions := []ds.Point{
		ds.NewPoint(1, 0),  // Forwards
		ds.NewPoint(0, 1),  // Vertical
		ds.NewPoint(1, 1),  // rigt diag
		ds.NewPoint(-1, 1), // left diag
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			start := ds.NewPoint(col, row)

			for _, dir := range directions {
				if check(grid, start, dir, target) {
					res += 1
				}

				if check(grid, start, dir, targetReverse) {
					res += 1
				}
			}

		}
	}
	return res
}

func solvePart2(grid [][]rune) int {
	res := 0
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[0])-1; col++ {
			if grid[row][col] != 'A' {
				continue
			}

			diag1 := false
			diag2 := false
			if (grid[row-1][col+1] == 'M' && grid[row+1][col-1] == 'S') || (grid[row-1][col+1] == 'S' && grid[row+1][col-1] == 'M') {
				diag1 = true
			}

			if (grid[row-1][col-1] == 'M' && grid[row+1][col+1] == 'S') || (grid[row-1][col-1] == 'S' && grid[row+1][col+1] == 'M') {
				diag2 = true
			}
			if diag1 && diag2 {
				res += 1
			}

		}
	}
	return res
}
