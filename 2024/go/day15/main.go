package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	ds "github.com/shadesfear/aoc-lib-go/datastructures"
	"github.com/shadesfear/aoc-lib-go/files"
)

type Point = ds.Point

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

func parse(input string) ([][]rune, string) {
	splt := strings.Split(input, "\n\n")

	grid := splt[0]
	i := splt[1]

	g, err := files.ParseInputToGrid(grid)
	if err != nil {
		log.Fatal(err)
	}
	return g, i
}

func start(grid [][]rune) ds.Point {
	for y, line := range grid {
		for x, char := range line {
			p := ds.NewPoint(x, y)
			if char == '@' {
				return p
			}
		}
	}
	log.Fatal("No start")
	panic("")
}

var dirs map[rune]ds.Point = map[rune]ds.Point{
	'^': ds.NewPoint(0, -1),
	'>': ds.NewPoint(1, 0),
	'v': ds.NewPoint(0, 1),
	'<': ds.NewPoint(-1, 0),
}

func gF(grid [][]rune, p ds.Point) rune {
	return grid[p.Y][p.X]
}

func move(grid [][]rune, point ds.Point, dir ds.Point) bool {
	next := point.Add(dir)
	curRune := gF(grid, point)

	if !point.InBounds(len(grid), len(grid[0])) {
		return false
	}

	if curRune == '#' {
		return false
	}

	if curRune == '.' {
		return true
	}

	if move(grid, next, dir) {
		grid[next.Y][next.X] = 'O'
		grid[point.Y][point.X] = '.'
		return true

	}
	return false
}

func do(grid [][]rune, cur ds.Point, instr rune) (bool, ds.Point) {
	dir := dirs[instr]
	next := cur.Add(dir)

	if move(grid, next, dir) {
		grid[cur.Y][cur.X] = '.'
		grid[next.Y][next.X] = '@'
		return true, next
	}

	return false, cur
}

func printGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println()
		for _, char := range line {
			fmt.Print(string(char))
		}
	}
}

func solvePart1(lines string) int {
	res := 0
	grid, instr := parse(lines)
	st := start(grid)

	posdir := []rune{'>', '<', 'v', '^'}

	for _, in := range instr {
		if !slices.Contains(posdir, in) {
			continue
		}
		_, st = do(grid, st, in)
	}

	for y, row := range grid {
		for x, char := range row {

			if char != 'O' {
				continue
			}

			res += 100*y + x
		}
	}

	return res
}

func toNewGrid(grid [][]rune) [][]rune {
	newGrid := [][]rune{}

	for _, row := range grid {
		newR := []rune{} // Should be 2x length of old row
		for _, char := range row {
			if char == '#' {
				newR = append(newR, '#')
				newR = append(newR, '#')
			}
			if char == 'O' {
				newR = append(newR, '[')
				newR = append(newR, ']')
			}
			if char == '.' {
				newR = append(newR, '.')
				newR = append(newR, '.')
			}
			if char == '@' {
				newR = append(newR, '@')
				newR = append(newR, '.')
			}

		}

		log.Println(newR)
		newGrid = append(newGrid, newR)
	}
	return newGrid
}

type Box struct {
	start, end Point
}

func getFullBox(grid [][]rune, pos ds.Point) (bool, Box) {
	if !pos.InBounds(len(grid), len(grid[0])) {
		return false, Box{}
	}
	if gF(grid, pos) == '[' {
		end := pos.Move(1, 0)

		if end.InBounds(len(grid), len(grid[0])) || gF(grid, end) == ']' {
			return true, Box{pos, end}
		}

	}

	if gF(grid, pos) == ']' {
		start := pos.Move(-1, 0)

		if start.InBounds(len(grid), len(grid[0])) || gF(grid, start) == '[' {
			return true, Box{pos, start}
		}

	}

	return false, Box{}
}

func findConnectedBoxes(grid [][]rune, startBox Box, dir Point) []Box {
	var boxes []Box
	boxes = append(boxes, startBox)

	if dir.X != 0 {
		y := startBox.start.Y
		for x := 0; x < len(grid[y]); x++ {
			if isBox, boxInfo := getFullBox(grid, ds.NewPoint(x, y)); isBox {
				alreadyFound := false
				for _, b := range boxes {
					if b.start == boxInfo.start {
						alreadyFound = true
						break
					}
				}

				if !alreadyFound {
					boxes = append(boxes, boxInfo)
				}
			}
		}
	} else if dir.Y != 0 {
		for y := 0; y < len(grid); y++ {
			for x := startBox.start.X; x <= startBox.end.X; x++ {
				if isBox, boxInfo := getFullBox(grid, ds.NewPoint(x, y)); isBox {
					alreadyFound := false
					for _, b := range boxes {
						if b.start == boxInfo.start {
							alreadyFound = true
							break
						}
					}
					if !alreadyFound {
						boxes = append(boxes, boxInfo)
					}
				}
			}
		}
	}

	return boxes
}

func move2(grid [][]rune, point ds.Point, dir ds.Point, visited map[Point]bool) bool {
	curRune := gF(grid, point)

	if !point.InBounds(len(grid), len(grid[0])) {
		return false
	}

	visited[point] = true

	if curRune == '#' {
		return false
	}

	if curRune == '.' {
		return true
	}

	isBox, boxInfo := getFullBox(grid, point)
	if !isBox {
		return false
	}

	boxes := findConnectedBoxes(grid, boxInfo, dir)

	for _, box := range boxes {
		newStart := box.start.Add(dir)
		newEnd := box.end.Add(dir)

		if !newStart.InBounds(len(grid), len(grid[0])) || newEnd.InBounds(len(grid), len(grid[0])) {
			return false
		}

		if gF(grid, newStart) != '.' && !visited[newStart] {
			if !move2(grid, newStart, dir, visited) {
				return false
			}
		}

		if gF(grid, newEnd) != '.' && !visited[newEnd] {
			if !move2(grid, newEnd, dir, visited) {
				return false
			}
		}

	}

	for _, box := range boxes {
		grid[box.start.Y][box.start.X] = '.'
		grid[box.end.Y][box.end.X] = '.'
	}

	for _, box := range boxes {
		newStart := box.start.Add(dir)
		newEnd := box.end.Add(dir)
		grid[newStart.Y][newStart.X] = '['
		grid[newEnd.Y][newEnd.X] = ']'
	}

	// if !newStart.InBounds(len(grid), len(grid[0])) && !newEnd.InBounds(len(grid), len(grid[0])) {
	// 	return false
	// }

	// if gF(grid, newStart) != '.' || gF(grid, newEnd) != '.' {
	// 	if gF(grid, newStart) == '[' {
	// 		if !move2(grid, newStart, dir) {
	// 			return false
	// 		}
	// 	} else {
	// 		return false
	// 	}
	// }

	// grid[boxStart.Y][boxStart.X] = '.'
	// grid[boxEnd.Y][boxEnd.X] = '.'
	// grid[newStart.Y][newStart.X] = '['
	// grid[newEnd.Y][newEnd.X] = ']'

	// if move2(grid, next, dir) {
	// 	grid[next.Y][next.X] = 'O'
	// 	grid[point.Y][point.X] = '.'
	// 	return true
	//
	// }
	return true
}

func do2(grid [][]rune, cur ds.Point, instr rune) (bool, ds.Point) {
	dir := dirs[instr]
	next := cur.Add(dir)

	if move2(grid, next, dir, map[Point]bool{}) {
		grid[cur.Y][cur.X] = '.'
		grid[next.Y][next.X] = '@'
		return true, next
	}

	return false, cur
}

func solvePart2(lines string) int {
	res := 0

	grid, _ := parse(lines)
	nGrid := toNewGrid(grid)
	st := start(nGrid)
	printGrid(nGrid)

	// posdir := []rune{'>', '<', 'v', '^'}

	_, st = do2(nGrid, st, '<')
	_, st = do2(nGrid, st, '<')
	_, st = do2(nGrid, st, '<')

	// for _, in := range instr {
	// 	if !slices.Contains(posdir, in) {
	// 		continue
	// 	}
	// 	_, st = do2(nGrid, st, in)
	// }

	printGrid(nGrid)

	for y, row := range grid {
		for x, char := range row {

			if char != '[' {
				continue
			}
			left := x
			right := len(row) - (x + 1)

			res += 100*y + (max(left, right))
		}
	}

	return res
}
