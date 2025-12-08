package main

import (
	"cmp"
	"log"
	"math"
	"slices"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/str"

	// "github.com/shadesfear/aoc-lib-go/math"

	"github.com/shadesfear/aoc-lib-go/datastructures"
)

type Point = datastructures.Point

type Point3d struct {
	x, y, z int
	circuit int
}

type PointPair struct {
	a, b Point3d
}

var DistanceMatrix = map[PointPair]float64{}

func (p Point3d) EuclideanDistance(other Point3d) float64 {
	return math.Sqrt(math.Pow(float64(p.x-other.x), 2) + math.Pow(float64(p.y-other.y), 2) + math.Pow(float64(p.z-other.z), 2))
}

func (p Point3d) Equal(other Point3d) bool {
	return p.x == other.x && p.y == other.y && p.z == other.z
}

func main() {
	lines, err := files.ReadInputLines("../../inputs/day08.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

func LinesToPoints(lines []string) []Point3d {
	var points = []Point3d{}

	for _, line := range lines {
		ns := strings.Split(line, ",")
		points = append(points, Point3d{
			x:       str.ToInt(ns[0]),
			y:       str.ToInt(ns[1]),
			z:       str.ToInt(ns[2]),
			circuit: -1,
		})
	}

	return points

}

func closest(points []Point3d, cand Point3d) (Point3d, int) {
	var close Point3d
	var idx int
	distance := math.MaxFloat64
	for i, p := range points {
		if p.Equal(cand) {
			continue
		}
		newDistance := p.EuclideanDistance(cand)
		if newDistance < distance {
			distance = newDistance
			close = p
			idx = i
		}
	}
	return close, idx
}

func part1(points []Point3d, connections int) int {
	maxCircuit := 0
	for range connections {
		for i := 0; i < len(points); i++ {
			if points[i].circuit != -1 {
				continue
			}
			close, idx := closest(points, points[i])

			if close.circuit != -1 {
				points[idx].circuit = close.circuit
			} else {
				points[i].circuit = maxCircuit
				points[idx].circuit = maxCircuit
				maxCircuit++
			}

		}
	}

	return maxCircuit
}

func minMap() PointPair {
	var minPair PointPair
	minDist := math.MaxFloat64

	for pair, dist := range DistanceMatrix {
		if dist < minDist {
			minDist = dist
			minPair = pair
		}
	}
	return minPair
}

func sortedPairs() []PointPair {
	pairs := make([]PointPair, 0, len(DistanceMatrix))
	for pair := range DistanceMatrix {
		pairs = append(pairs, pair)
	}

	slices.SortFunc(pairs, func(a, b PointPair) int {
		return cmp.Compare(DistanceMatrix[a], DistanceMatrix[b])
	})
	return pairs
}

func solvePart1(lines []string) int {

	points := LinesToPoints(lines)
	log.Println(len(points))

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			ppair := PointPair{
				a: points[i],
				b: points[j],
			}
			distance := ppair.a.EuclideanDistance(ppair.b)

			DistanceMatrix[ppair] = distance
		}

	}

	log.Println(len(DistanceMatrix))
	pairs := sortedPairs()

	for _, pair := range pairs[:10]{
		if pair.a.circuit == -1 && pair.b.circuit == -1 {

		}
	}

	res := part1(points, 10)

	return res
}

func solvePart2(lines []string) int {
	res := 0

	return res
}
