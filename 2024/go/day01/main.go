package main

import (
	"log"
	"sort"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/math"
)

func main() {
	lines, err := files.ReadInputLines("input.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

func parsePart1(lines []string) ([]int, []int) {
	left, right := []int{}, []int{}

	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		ints, err := files.ParseInputToInts(line)
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, ints[0])
		right = append(right, ints[1])
	}
	return left, right
}

func parsePart2(lines []string) ([]int, map[int]int) {
	left, right := parsePart1(lines)
	counts := map[int]int{}

	for _, i := range right {
		counts[i] += 1
	}
	return left, counts
}

func solvePart1(lines []string) int {
	left, right := parsePart1(lines)
	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	res := 0

	for i := 0; i < len(left); i++ {
		res += math.Abs(right[i] - left[i])
	}

	return res
}

func solvePart2(lines []string) int {
	res := 0
	left, counts := parsePart2(lines)
	for _, i := range left {
		c := counts[i]
		res += i * c
	}
	return res
}
