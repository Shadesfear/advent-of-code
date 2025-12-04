package main

import (
	"log"
	"strings"

	"github.com/shadesfear/aoc-lib-go/datastructures"
	"github.com/shadesfear/aoc-lib-go/files"
)

type Point = datastructures.Point

func main() {
	lines, err := files.ReadInputLines("../../inputs/day03.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

func findMax(list string, start, end int) (int, int) {
	// log.Printf("List: %s, start: %d end: %d", list, start, end)
	var maxJolts uint8
	var idx int
	for i := start; i < end; i++ {
		if list[i] > maxJolts {
			maxJolts = list[i]
			idx = i
		}
	}
	return int(maxJolts - '0'), idx
}

func findJolts(bank string, digits int) int {
	result := 0
	idx := -1
	for i := digits; i > 0; i-- {
		var jolts int
		jolts, idx = findMax(bank, idx+1, len(bank)-i+1)
		result = result*10 + jolts
	}
	return result

}

func solvePart1(lines []string) int {
	res := 0

	for _, line := range lines {
		line = strings.ReplaceAll(line, "\n", "")
		res += findJolts(line, 2)
	}

	return res
}

func solvePart2(lines []string) int {
	res := 0

	for _, line := range lines {

		line = strings.ReplaceAll(line, "\n", "")
		res += findJolts(line, 12)
	}

	return res
}
