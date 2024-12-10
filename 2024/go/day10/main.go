package main

import (
	"log"

	ds "github.com/shadesfear/aoc-lib-go/datastructures"
	"github.com/shadesfear/aoc-lib-go/files"
)

func main() {
	lines, err := files.ReadInputFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}

	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

func parse(input string) [][]int {
	intGrid := [][]int{}
	lines := files.StringToLines(input)

	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		ro := []int{}
		for _, char := range line {
			if char == '.' {
				ro = append(ro, -1)
				continue
			}
			ro = append(ro, int(char-'0'))
		}
		intGrid = append(intGrid, ro)

	}

	return intGrid
}

func findZeroes(inp [][]int) []ds.Point {
	res := []ds.Point{}
	for y, row := range inp {
		for x, col := range row {
			if col != 0 {
				continue
			}
			res = append(res, ds.NewPoint(x, y))
		}
	}
	return res
}

func traverse(grid [][]int, start ds.Point, seen map[ds.Point]bool) (int, int) {
	res := 0
	for _, n := range start.Neighbours() {

		if !n.InBounds(len(grid), len(grid[0])) {
			continue
		}
		if !(grid[n.Y][n.X]-grid[start.Y][start.X] == 1) {
			continue
		}

		if grid[n.Y][n.X] == 9 {
			res++
			seen[n] = true
			continue
		}

		r, _ := traverse(grid, n, seen)
		res += r
	}

	return res, len(seen)
}

func solvePart1(input string) int {
	grid := parse(input)
	zeroes := findZeroes(grid)

	res := 0

	for _, zero := range zeroes {
		_, r := traverse(grid, zero, map[ds.Point]bool{})
		res += r
	}

	return res
}

func solvePart2(input string) int {
	grid := parse(input)
	zeroes := findZeroes(grid)

	res := 0
	for _, zero := range zeroes {
		r, _ := traverse(grid, zero, map[ds.Point]bool{})
		res += r
	}

	return res
}
