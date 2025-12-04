package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/str"

	// "github.com/shadesfear/aoc-lib-go/math"

	"github.com/shadesfear/aoc-lib-go/datastructures"
)

type Point = datastructures.Point

var OFFSETS [][]int = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func main() {
	input, err := files.ReadDayInput(4)
	if err != nil {
		log.Fatal("Cant read input file")
	}

	grid, err := files.ParseInputToGrid(input)
	if err != nil {
		log.Fatal("cant parse to grid")
	}
	// part1 := solvePart1(grid)
	// log.Println(part1)

	// part2 := solvePart2(grid)
	// log.Println(part2)

	_, hist := solvePart2CaptureGrid(grid)
	for _, g := range hist {
		fmt.Print("\033[H\033[2J")

		// fmt.Print("\033[H") // move to top-left, no clear
		str.PrettyPrintGrid(g)
		time.Sleep(300 * time.Millisecond)
	}
}

func rollsOfPaper(grid [][]rune, x, y int) int {
	counter := 0
	rows, cols := len(grid), len(grid[0])
	for _, dir := range OFFSETS {
		nX := x + dir[0]
		nY := y + dir[1]
		if nX < 0 || nX >= cols || nY < 0 || nY >= rows {
			continue
		}
		if grid[nY][nX] == '@' {
			counter++
		}

	}

	// if grid[y][x] == '@' {
	// 	grid[y][x] = rune(counter)
	// }
	return counter
}

func solvePart1(grid [][]rune) int {
	res := 0
	rows, cols := len(grid), len(grid[0])

	countGrid := make([][]int, rows)
	for i := range cols {
		countGrid[i] = make([]int, rows)
	}

	// str.PrettyPrintGrid(grid)

	for y := range rows {
		for x := range cols {
			if grid[y][x] == '.' {
				continue
			}

			nRolls := rollsOfPaper(grid, x, y)
			if nRolls < 4 {
				res++
				countGrid[y][x] = nRolls
			}

		}
	}

	// str.PrettyPrintGrid(countGrid)

	return res
}

func solvePart2(grid [][]rune) int {
	res := 0

	rows, cols := len(grid), len(grid[0])

	for {
		var toDel = [][]int{}

		for y := range rows {
			for x := range cols {
				if grid[y][x] == '.' {
					continue
				}

				nRolls := rollsOfPaper(grid, x, y)
				if nRolls < 4 {
					toDel = append(toDel, []int{x, y})
				}

			}
		}

		toDelCount := len(toDel)
		if toDelCount == 0 {
			break
		}

		res += toDelCount

		for _, coord := range toDel {
			grid[coord[1]][coord[0]] = '.'
		}

	}

	return res
}

func copyGrid[T any](grid [][]T) [][]T {
	newGrid := make([][]T, len(grid))
	for i := range grid {
		newGrid[i] = make([]T, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

func solvePart2CaptureGrid(grid [][]rune) (int, [][][]rune) {
	res := 0

	rows, cols := len(grid), len(grid[0])

	historic := [][][]rune{}

	for {
		var toDel = [][]int{}
		newGrid := copyGrid(grid)
		historic = append(historic, newGrid)

		for y := range rows {
			for x := range cols {
				if grid[y][x] == '.' {
					continue
				}

				nRolls := rollsOfPaper(grid, x, y)
				if nRolls < 4 {
					toDel = append(toDel, []int{x, y})
				}

			}
		}

		toDelCount := len(toDel)
		if toDelCount == 0 {
			break
		}

		res += toDelCount

		for _, coord := range toDel {
			grid[coord[1]][coord[0]] = '.'
		}

	}

	return res, historic
}
