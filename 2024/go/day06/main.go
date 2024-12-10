package main

import (
	"log"

	ds "github.com/shadesfear/aoc-lib-go/datastructures"
	"github.com/shadesfear/aoc-lib-go/files"
	// "github.com/shadesfear/aoc-lib-go/str"
	// "github.com/shadesfear/aoc-lib-go/math"
)

var (
	UP    ds.Point = ds.NewPoint(0, -1)
	DOWN  ds.Point = ds.NewPoint(0, 1)
	LEFT  ds.Point = ds.NewPoint(-1, 0)
	RIGHT ds.Point = ds.NewPoint(1, 0)
)

var (
	visited  map[ds.Point]bool
	GRID     [][]rune
	ROWS     int
	COLS     int
	START    ds.Point
	STARTDIR ds.Point
	DIRS     []ds.Point = []ds.Point{UP, RIGHT, DOWN, LEFT}
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

func findStartPos(grid [][]rune) {
	s := ds.NewPoint(0, 0)
	for y, row := range grid {
		for x, char := range row {
			if char == '^' {
				s = ds.NewPoint(x, y)
			} else if char == '>' {
				s = ds.NewPoint(x, y)
			} else if char == 'v' {
				s = ds.NewPoint(x, y)
			} else if char == '<' {
				s = ds.NewPoint(x, y)
			}
		}
	}
	START = s
	STARTDIR = UP
}

// func rotateDir(dir int) ds.Point {
// 	// rotationMap := map[ds.Point]ds.Point{
// 	// 	UP:    RIGHT,
// 	// 	RIGHT: DOWN,
// 	// 	DOWN:  LEFT,
// 	// 	LEFT:  UP,
// 	// }
// 	//
// 	// if newDir, exists := rotationMap[dir]; exists {
// 	// 	return newDir
// 	// }
//
// 	fmt.Printf("Warning: Invalid direction %v\n", dir)
// 	return dir
// }

func traverse(start, startDir ds.Point, grid [][]rune) map[ds.Point]bool {
	rows, cols := len(grid), len(grid[0])
	visited := map[ds.Point]bool{}
	current := start
	dirC := 0

	for {
		visited[current] = true
		dir := DIRS[dirC%4]
		next := current.Move(dir.X, dir.Y)

		if !next.InBounds(rows, cols) {
			break
		}

		if grid[next.Y][next.X] == '#' {
			dirC++
			continue
		}
		current = next

	}
	return visited
}

func solvePart1(lines string) int {
	grid, err := files.ParseInputToGrid(lines)
	if err != nil {
		log.Fatal(err)
	}

	GRID = grid

	ROWS = len(grid)
	COLS = len(grid[0])

	findStartPos(grid)

	visited = traverse(START, STARTDIR, grid)

	return len(visited)
}

func loopDetector(start, dir ds.Point, obstacle ds.Point) bool {
	cur := start
	newDirC := 0

	visited := map[int64]bool{}

	for {

		d := DIRS[newDirC%4]
		key := int64(cur.X) | (int64(cur.Y) << 16) | (int64(d.X+1) << 32) | (int64(d.Y) << 34)

		if visited[key] {
			return true
		}

		visited[key] = true
		next := cur.Move(d.X, d.Y)

		if !next.InBounds(ROWS, COLS) {
			break
		}

		if GRID[next.Y][next.X] == '#' || (next == obstacle) {
			newDirC++
			continue
		}

		cur = next

	}

	return false
}

func solvePart2(lines string) int {
	res := 0

	for pos := range visited {
		loop := loopDetector(START, STARTDIR, pos)
		if loop {
			res++
		}
	}

	return res
}
