package main

import (
	"log"
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

func insidePolygon(vertices []Point, point Point) bool {
	crossings := 0
	n := len(vertices)
	for _, v := range vertices {
		if v == point {
			return true
		}
	}
	for i := range n {
		a := vertices[i]
		b := vertices[(i+1)%n]

		if a.Y == b.Y {
			continue
		}

		if a.Y > b.Y {
			a, b = b, a
		}

		if point.Y >= a.Y && point.Y < b.Y {
			if a.X > point.X {
				crossings++
			}
		}
	}

	return crossings%2 == 1

}

func segmentsIntersects(h1, h2, v1, v2 Point) bool {
	hY := h1.Y
	vX := v1.X

	hxMin, hxMax := min(h1.X, h2.X), max(h1.X, h2.X)
	vyMin, vyMax := min(v1.Y, v2.Y), max(v1.Y, v2.Y)

	return vX > hxMin && vX < hxMax && hY > vyMin && hY < vyMax
}

func rectangleCrossesPolygon(a, b Point, vertices []Point) bool {
	minX, maxX := min(a.X, b.X), max(a.X, b.X)
	minY, maxY := min(a.Y, b.Y), max(a.Y, b.Y)

	top := [2]Point{{X: minX, Y: minY}, {X: maxX, Y: minY}}
	bot := [2]Point{{X: minX, Y: maxY}, {X: maxX, Y: maxY}}
	left := [2]Point{{X: minX, Y: minY}, {X: minX, Y: maxY}}
	right := [2]Point{{X: maxX, Y: minY}, {X: maxX, Y: maxY}}

	n := len(vertices)

	for i := range n {
		p1 := vertices[i]
		p2 := vertices[(i+1)%n]

		if p1.Y == p2.Y {
			if segmentsIntersects(p1, p2, left[0], left[1]) {
				return true
			}
			if segmentsIntersects(p1, p2, right[0], right[1]) {
				return true
			}
		} else {
			if segmentsIntersects(top[0], top[1], p1, p2) {
				return true
			}
			if segmentsIntersects(bot[0], bot[1], p1, p2) {
				return true
			}
		}
	}
	return false
}

func part2(points []Point) int {
	var biggest int
	var aa Point
	var bb Point
	n := len(points)

	for i := range n {
		for j := i + 1; j < n; j++ {
			a, b := points[i], points[j]
			corners := Corners(a, b)
			insideb := true
			for _, corner := range corners {
				if !insidePolygon(points, corner) {
					insideb = false
					break
				}
			}

			if insideb {
				if rectangleCrossesPolygon(a, b, points) {
					continue
				}
				dx := mymath.Abs(b.X - a.X)
				dy := mymath.Abs(b.Y - a.Y)
				area := (dx + 1) * (dy + 1)
				if area >= biggest {
					biggest = area
					aa = a
					bb = b
				}
			}

		}
	}
	log.Println(aa, bb)
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
