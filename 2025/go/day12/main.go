package main

import (
	"log"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/str"

	// "github.com/shadesfear/aoc-lib-go/math"

	"github.com/shadesfear/aoc-lib-go/datastructures"
)

type Point = datastructures.Point

type Shape struct {
	width  int
	points []Point
}

func (s Shape) RotateCW() Shape {
	var points []Point
	for _, point := range s.points {
		newX := 1 - (point.Y - 1)
		newY := 1 + (point.X - 1)
		points = append(points, Point{
			X: newX,
			Y: newY,
		})
	}
	return Shape{width: s.width, points: points}
}

func (s Shape) Print() {
	grid := [][]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', '.', '.'},
	}
	for _, point := range s.points {
		grid[point.Y][point.X] = '#'
	}

	str.PrettyPrintGrid(grid)
}

type ShapeLocation struct {
	shape  Shape
	center Point
}

type Grid struct {
	Shapes      []Shape
	rows        int
	cols        int
	Amounts     []int
	AddedShapes []ShapeLocation
}

func (g *Grid) AddShape(shape int, center Point) {
	g.AddedShapes = append(g.AddedShapes, ShapeLocation{shape: g.Shapes[shape], center: center})
}
func (g Grid) Print() {
	grid := [][]rune{}
	for row := 0; row < g.rows; row++ {
		grid = append(grid, []rune{})
		for cols := 0; cols < g.rows; cols++ {
			grid[row] = append(grid[row], '.')
		}
	}

	for _, shape := range g.AddedShapes {
		start := shape.center.Move(-1, -1)
		for _, point := range shape.shape.points {
			grid[start.Y+point.Y][start.X+point.X] = '#'
		}
	}
	str.PrettyPrintGrid(grid)
}

func parse(input string) []Grid {
	parts := strings.Split(input, "\n\n")
	shapes := []Shape{}
	for _, str := range parts[:len(parts)-1] {
		shapeLines := strings.Split(str, "\n")
		points := []Point{}

		for i, shapeLine := range shapeLines[1:] {
			for j := 0; j < len(shapeLine); j++ {
				if shapeLine[j] == '#' {
					points = append(points, Point{X: j, Y: i})
				}
			}
		}
		shapes = append(shapes, Shape{width: 3, points: points})
	}

	grids := []Grid{}
	for line := range strings.SplitSeq(parts[len(parts)-1], "\n") {
		if line == "" {
			continue
		}
		puzzle := strings.Split(line, ": ")
		gridS := strings.Split(puzzle[0], "x")
		rows, cols := str.ToInt(gridS[0]), str.ToInt(gridS[1])

		numbers, err := files.ParseInputToInts(puzzle[1])
		if err != nil {
			panic("No ints :(")
		}

		grid := Grid{
			Shapes:  shapes,
			rows:    rows,
			cols:    cols,
			Amounts: numbers,
		}

		grids = append(grids, grid)

	}

	return grids
}

func solvePart1(lines string) int {
	res := 0
	grids := parse(lines)
	for _, grid := range grids {
		totalShapes := 0
		for _, n := range grid.Amounts {
			totalShapes += n
		}
		fits := (grid.rows / 3) * (grid.cols / 3)
		if fits >= totalShapes {
			res++
		} else {
			totalArea := 0
			for i, n := range grid.Amounts {
				totalArea += n * len(grid.Shapes[i].points)
			}
		}
	}
	// grids[0].AddShape(0, Point{X: 1, Y: 1})
	// grids[0].Print()

	return res
}

func solvePart2(lines string) int {
	res := 0

	return res
}

func main() {
	str, err := files.ReadInputFile("../../inputs/day12.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(str)
	log.Println(part1)

	part2 := solvePart2(str)
	log.Println(part2)
}
