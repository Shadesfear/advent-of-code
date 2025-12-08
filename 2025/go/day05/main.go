package main

import (
	"log"
	"sort"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/str"

	// "github.com/shadesfear/aoc-lib-go/math"

	"github.com/shadesfear/aoc-lib-go/datastructures"
)

type Point = datastructures.Point

func main() {
	// lines, err := files.ReadInputLines("../../inputs/day05.txt")
	input, err := files.ReadDayInput(5)
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(input)
	log.Println(part1)

	part2 := solvePart2(input)
	log.Println(part2)
}

type Range struct {
	left  int
	right int
}

type Database struct {
	Ranges []Range
	IDs    []int
}

func RangeFromLine(line string) Range {
	s := strings.Split(line, "-")
	leftString := s[0]
	rightString := s[1]
	left := str.ToInt(leftString)
	right := str.ToInt(rightString)
	return Range{left, right}
}

func Parse(lines string) Database {
	split := strings.Split(lines, "\n\n")

	var ranges []Range
	for _, line := range strings.Split(split[0], "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "-")
		ranges = append(ranges, Range{str.ToInt(parts[0]), str.ToInt(parts[1])})
	}

	var ids []int
	for _, line := range strings.Split(split[1], "\n") {
		if line == "" {
			continue
		}
		ids = append(ids, str.ToInt(line))
	}

	return Database{ranges, ids}
}

func IDInRange(id int, r Range) bool {
	if id < r.left || id > r.right {
		return false
	}
	return true
}

func solvePart1(lines string) int {
	res := 0
	db := Parse(lines)

	for _, id := range db.IDs {
		for _, r := range db.Ranges {
			if IDInRange(id, r) {
				res++
				break
			}
		}
	}

	return res
}

func SolvePart2Parsed(db Database) int {
	res := 0

	ranges := db.Ranges
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].left < ranges[j].left
	})

	newRange := []Range{}
	newRange = append(newRange, ranges[0])

	for i := 1; i < len(ranges); i++ {
		last := &newRange[len(newRange)-1]
		curr := ranges[i]

		if curr.left <= last.right {
			last.right = max(last.right, curr.right)
		} else {
			newRange = append(newRange, curr)
		}

	}

	for _, nr := range newRange {
		res += (nr.right - nr.left) + 1
	}
	return res

}

func solvePart2(lines string) int {
	db := Parse(lines)
	return SolvePart2Parsed(db)
}
