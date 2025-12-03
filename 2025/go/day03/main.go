package main

import (
	"log"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/str"

	// "github.com/shadesfear/aoc-lib-go/str"

	"github.com/shadesfear/aoc-lib-go/math"

	"github.com/shadesfear/aoc-lib-go/datastructures"
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

func PowerBankLargest(s string) int {

	largest := 0
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			si := string(s[i])
			sj := string(s[j])
			ns := si + sj
			n := str.ToInt(ns)
			if n > largest {
				largest = n
			}

		}
	}

	return largest
}

func recurse(s string, n int) int {
	if n == 0 {
		return 0
	}

	if len(s) == n {
		return str.ToInt(s)
	}

	n1 := str.ToInt(s[0:1])

	a := n1*(math.Pow(10, n-1)) + recurse(s[1:], n-1)
	b := recurse(s[1:], n)
	return max(a, b)
}

func PowerBankLargest12(s string) int {

	largest := 0

	return largest
}

func solvePart1(lines []string) int {
	res := 0

	for _, line := range lines {
		line = strings.ReplaceAll(line, "\n", "")
		res += PowerBankLargest(line)
	}

	return res
}

func solvePart2(lines []string) int {
	res := 0

	for _, line := range lines {

		line = strings.ReplaceAll(line, "\n", "")
		res += recurse(line, 12)
	}

	return res
}
