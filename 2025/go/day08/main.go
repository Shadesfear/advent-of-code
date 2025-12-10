package main

import (
	"cmp"
	"log"
	"math"
	"slices"
	"sort"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/str"
	// "github.com/shadesfear/aoc-lib-go/math"
	// "github.com/shadesfear/aoc-lib-go/datastructures"
)

type JBox struct {
	x, y, z int
	parent  int
}

type PointPair struct {
	a, b int
}

var DistanceMatrixPoint = map[PointPair]float64{}

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

	for i, line := range lines {
		ns := strings.Split(line, ",")
		points = append(points, JBox{
			x:      str.ToInt(ns[0]),
			y:      str.ToInt(ns[1]),
			z:      str.ToInt(ns[2]),
			parent: i,
		})
	}

	return points

}

func createDistanceMapJBox(JBoxes []JBox) {
	for i := range JBoxes {
		for j := i + 1; j < len(JBoxes); j++ {
			pair := PointPair{
				a: i,
				b: j,
			}

			DistanceMatrixPoint[pair] = JBoxes[i].EuclideanDistance(JBoxes[j])
		}

	}
}

func find(boxes []JBox, i int) int {
	for boxes[i].parent != i {
		i = boxes[i].parent
	}
	return i
}

func Union(boxes []JBox, a, b int) {
	rootA := find(boxes, a)
	rootB := find(boxes, b)
	if rootA != rootB {
		boxes[rootA].parent = rootB
	}
}

func sortedPairsPoint() []PointPair {
	pairs := make([]PointPair, 0, len(DistanceMatrixPoint))
	for pair := range DistanceMatrixPoint {
		pairs = append(pairs, pair)
	}

	slices.SortFunc(pairs, func(a, b PointPair) int {
		return cmp.Compare(DistanceMatrixPoint[a], DistanceMatrixPoint[b])
	})
	return pairs
}

func prod(xs []int) int {
	res := xs[0]
	for i := 1; i < len(xs); i++ {
		res *= xs[i]
	}
	return res
}

func part1(boxes []JBox, n int) int {
	pairs := sortedPairsPoint()

	for _, pair := range pairs[:n] {
		Union(boxes, pair.a, pair.b)
	}

	counts := map[int]int{}
	for _, box := range boxes {
		parent := find(boxes, box.parent)
		if _, ok := counts[parent]; !ok {
			counts[parent] = 0
		}
		counts[parent]++
	}

	c := []int{}

	for val := range counts {
		c = append(c, counts[val])
	}

	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})

	return prod(c[:3])

}

func part2(boxes []JBox) int {
	pairs := sortedPairsPoint()

	var res int
	for _, pair := range pairs {

		Union(boxes, pair.a, pair.b)
		parent := find(boxes, boxes[0].parent)
		allTheSame := true
		for _, box := range boxes {
			boxParent := find(boxes, box.parent)
			if boxParent != parent {
				allTheSame = false
			}
		}

		if allTheSame {
			res = boxes[pair.a].x * boxes[pair.b].x
			break
		}

	}

	return res
}

func solvePart1(lines []string) int {

	boxes := LinesToPoints(lines)
	createDistanceMapJBox(boxes)
	return part1(boxes, 1000)
}

func solvePart2(lines []string) int {
	boxes := LinesToPoints(lines)
	createDistanceMapJBox(boxes)

	return part2(boxes)
}

func findCircuitNaive(circuits [][]int, box int) int {
	for i, circuit := range circuits {
		for _, b := range circuit {
			if b == box {
				return i
			}
		}
	}
	return -1
}

func part1Naive(boxes []JBox, n int) int {
	pairs := sortedPairsPoint()

	circuits := make([][]int, len(boxes))
	for i := range boxes {
		circuits[i] = []int{i}
	}

	for _, pair := range pairs[:n] {
		circuitA := findCircuitNaive(circuits, pair.a)
		circuitB := findCircuitNaive(circuits, pair.b)

		if circuitA != circuitB {
			circuits[circuitA] = append(circuits[circuitA], circuits[circuitB]...)
			circuits[circuitB] = []int{}
		}
	}

	sizes := []int{}
	for _, circuit := range circuits {
		if len(circuit) > 0 {
			sizes = append(sizes, len(circuit))
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return prod(sizes[:3])
}
