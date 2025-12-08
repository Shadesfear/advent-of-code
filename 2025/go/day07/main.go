package main

import (
	"log"

	"github.com/shadesfear/aoc-lib-go/files"

	// "github.com/shadesfear/aoc-lib-go/math"

	"github.com/shadesfear/aoc-lib-go/datastructures"
)

type Point = datastructures.Point

var Down = datastructures.Down
var Left = datastructures.Left
var Right = datastructures.Right

func main() {
	lines, err := files.ReadInputLines("../../inputs/day07.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

func part1(grid *datastructures.Grid[rune]) int {
	splits := 0
	beams := []Point{
		{X: grid.Cols / 2, Y: 0},
	}

	for len(beams) > 0 {
		newBeams := []Point{}
		for _, beam := range beams {
			movedBeam := beam.MoveDir(datastructures.Down)
			if !grid.InBounds(movedBeam) {
				continue
			}
			if grid.Get(movedBeam) == '^' {
				seenLeft, seenRight := false, false
				newLeft := movedBeam.MoveDir(Left)
				newRight := movedBeam.MoveDir(Right)

				for _, nb := range newBeams {
					if nb.Equal(newLeft) {
						seenLeft = true
					}
					if nb.Equal(newRight) {
						seenRight = true
					}
				}
				if !seenLeft && grid.Get(newLeft) != '^' {
					newBeams = append(newBeams, newLeft)
					grid.Set(newLeft, '|')

				}
				if !seenRight && grid.Get(newRight) != '^' {
					newBeams = append(newBeams, newRight)
					grid.Set(newRight, '|')
				}

				// grid.Print()
				splits++
				// if splits > 15 {
				// 	log.Println("")
				// }
			} else {
				seen := false
				for _, nb := range newBeams {
					if movedBeam.Equal(nb) {
						seen = true
					}
				}
				if !seen {
					grid.Set(movedBeam, '|')
					newBeams = append(newBeams, movedBeam)
				}
			}
			// grid.Print()

		}

		beams = newBeams
	}
	return splits
}

func part2(grid *datastructures.Grid[rune]) int {
	timelines := 0

	beams := datastructures.NewStack[Point]()
	beams.Push(Point{X: grid.Cols / 2, Y: 0})
	memo := map[Point]int{}

	for !beams.IsEmpty() {
		cur, ok := beams.Pop()
		if !ok {
			panic("WRong pop")
		}
		count, ok := memo[cur]
		if ok {
			timelines += count
			continue
		}

		moved := cur.MoveDir(Down)
		if !grid.InBounds(moved) {
			timelines++
			continue
		}

		if grid.Get(moved) == '^' {
			leftBeam := moved.MoveDir(Left)
			rightBeam := moved.MoveDir(Right)
			beams.Push(leftBeam)
			beams.Push(rightBeam)
		} else {
			beams.Push(moved)
		}

	}

	return timelines
}

func solvePart1(lines []string) int {
	runes, err := files.ParseLinesToGrid(lines)
	if err != nil {
		panic("Could not make grid")
	}
	grid := datastructures.NewGrid(runes)

	res := part1(grid)
	return res
}

func solvePart2(lines []string) int {
	runes, err := files.ParseLinesToGrid(lines)
	if err != nil {
		panic("Could not make grid")
	}
	grid := datastructures.NewGrid(runes)

	res := part2(grid)
	return res
}
