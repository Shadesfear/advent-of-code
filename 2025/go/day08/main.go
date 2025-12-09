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

type JBox struct {
	x, y, z int
	circuit int
}

type Circuit struct {
	JBoxes []JBox
}

func (c Circuit) EuDist(other Circuit) float64 {
	minDist := math.MaxFloat64
	for _, jbox := range c.JBoxes {
		for _, otherbox := range other.JBoxes {
			dist := jbox.EuclideanDistance(otherbox)
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

type PointPair struct {
	a, b JBox
}

type CircuitPair struct {
	a, b Circuit
}

var DistanceMatrix = map[*CircuitPair]float64{}

func (p JBox) EuclideanDistance(other JBox) float64 {
	return math.Sqrt(math.Pow(float64(p.x-other.x), 2) + math.Pow(float64(p.y-other.y), 2) + math.Pow(float64(p.z-other.z), 2))
}

func (p JBox) Equal(other JBox) bool {
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

func LinesToPoints(lines []string) []JBox {
	var points = []JBox{}

	for _, line := range lines {
		ns := strings.Split(line, ",")
		points = append(points, JBox{
			x:       str.ToInt(ns[0]),
			y:       str.ToInt(ns[1]),
			z:       str.ToInt(ns[2]),
			circuit: -1,
		})
	}

	return points

}

func closest(points []JBox, cand JBox) (JBox, int) {
	var close JBox
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

func part1(points []JBox, connections int) int {
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

func createDistanceMap(circuits []Circuit) {
	for i := range circuits {
		for j := i + 1; j < len(circuits); j++ {
			a, b := circuits[i], circuits[j]
			distance := a.EuDist(b)
			cpair := CircuitPair{
				a: a,
				b: b,
			}

			DistanceMatrix[&cpair] = distance
		}

	}

}

func initCircuits(points []JBox) []Circuit {
	var circuits = make([]Circuit, len(points))
	for _, point := range points {
		circuits = append(circuits, Circuit{
			JBoxes: []JBox{
				point,
			},
		})
	}
	return circuits
}

func solvePart1(lines []string) int {

	points := LinesToPoints(lines)
	circuits := initCircuits(points)
	createDistanceMap(circuits)
	pairs := sortedPairs()

	for _, pair := range pairs[:10] {
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
