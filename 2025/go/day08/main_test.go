package main

import (
	"testing"

	"github.com/shadesfear/aoc-lib-go/files"
)

func TestPart1(t *testing.T) {
	input := []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}
	exp := 40

	boxes := LinesToPoints(input)
	createDistanceMapJBox(boxes)

	res := part1(boxes, 10)
	if res != exp {
		t.Errorf("solvePart1() = %d, exp %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}
	exp := 25272

	DistanceMatrixPoint = map[PointPair]float64{}
	boxes := LinesToPoints(input)
	createDistanceMapJBox(boxes)

	res := part2(boxes)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}

var benchInput []string
var benchBoxes []JBox

func init() {
	benchInput, _ = files.ReadInputLines("../../inputs/day08.txt")
	benchBoxes = LinesToPoints(benchInput)
	createDistanceMapJBox(benchBoxes)
}

func resetBoxes() []JBox {
	boxes := make([]JBox, len(benchBoxes))
	copy(boxes, benchBoxes)
	for i := range boxes {
		boxes[i].parent = i
	}
	return boxes
}

func BenchmarkPart1UnionFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boxes := resetBoxes()
		part1(boxes, 1000)
	}
}

func BenchmarkPart1Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boxes := resetBoxes()
		part1Naive(boxes, 1000)
	}
}
