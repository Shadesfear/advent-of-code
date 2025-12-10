package main

import (
	"log"
	"math"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"
	mymath "github.com/shadesfear/aoc-lib-go/math"
	"github.com/shadesfear/aoc-lib-go/str"

	// "github.com/shadesfear/aoc-lib-go/math"

	"github.com/shadesfear/aoc-lib-go/datastructures"
)

type Point = datastructures.Point

func main() {
	lines, err := files.ReadInputLines("../../inputs/day09.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

func LinesToPoints(lines []string) []Point {
	var points []Point
	for _, line := range lines {
		split := strings.Split(line, ",")
		point := Point{
			X: str.ToInt(split[0]),
			Y: str.ToInt(split[1]),
		}
		points = append(points, point)
	}
	return points
}

func part1(points []Point) int {
	var biggest int
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := mymath.Abs(points[j].X - points[i].X)
			dy := mymath.Abs(points[j].Y - points[i].Y)
			area := (dx + 1) * (dy + 1)
			if area >= biggest {
				biggest = area
			}

		}
	}
	return biggest

}

func Corners(a, b Point) []Point {
	return []Point{
		a, b, {X: a.X, Y: b.Y}, {X: b.X, Y: a.Y},
	}
}

func inside(boundary []Point, p Point, minx, maxx, miny, maxy int) bool {
	boundarySet := datastructures.NewSet[Point]()
	boundarySet.AddM(boundary)

	if boundarySet.Contains(p) {
		return true
	}

	crossings := 0
	cur := p
	for cur.InBoundsA(minx, maxx, miny, maxy) {
		if boundarySet.Contains(cur) {
			crossings++
		}
		cur = cur.Move(1, 1)
	}
	if crossings%2 == 0 {
		// Not inside
		return false
	}
	return true
}

func insidePolygon(args) return type{

}

func part2(points []Point) int {
	var biggest int
	n := len(points)
	// First find green or red points on boundary
	boundary := []Point{}

	for i := 0; i < n; i++ {

		next := points[(i+1)%n]
		if next.X == points[i].X {
			for j := min(points[i].Y, next.Y); j < max(points[i].Y, next.Y); j++ {
				boundary = append(boundary, Point{X: points[i].X, Y: j})
			}
		} else {
			for j := min(points[i].X, next.X); j < max(points[i].X, next.X); j++ {
				boundary = append(boundary, Point{X: j, Y: points[i].Y})
			}
		}
	}

	var (
		minx = math.MaxInt
		maxx = 0
		miny = math.MaxInt
		maxy = 0
	)

	for _, point := range boundary {
		if point.X < minx {
			minx = point.X - 1
		}
		if point.X > maxx {
			maxx = point.X + 1
		}
		if point.Y < miny {
			miny = point.Y - 1
		}
		if point.Y > maxy {
			maxy = point.Y + 1
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a, b := points[i], points[j]
			corners := Corners(a, b)
			insideb := true
			for _, corner := range corners {
				if !inside(boundary, corner, minx, maxx, miny, maxy) {
					insideb = false
					break
				}
			}

			if insideb {
				dx := mymath.Abs(b.X - a.X)
				dy := mymath.Abs(b.Y - a.Y)
				area := (dx + 1) * (dy + 1)
				if area >= biggest {
					biggest = area
				}
			}

		}
	}
	return biggest
}

func solvePart1(lines []string) int {
	res := 0
	points := LinesToPoints(lines)
	res = part1(points)

	return res
}

func solvePart2(lines []string) int {
	res := 0
	points := LinesToPoints(lines)
	res = part2(points)
	return res
}
