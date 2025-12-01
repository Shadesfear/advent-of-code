package main

import (
	"github.com/shadesfear/aoc-lib-go/files"
	"log"
	// "github.com/shadesfear/aoc-lib-go/math"

	"github.com/shadesfear/aoc-lib-go/datastructures"
	"github.com/shadesfear/aoc-lib-go/str"
)

type Point = datastructures.Point

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

func solvePart1(lines []string) int {
	res := 0

	turningTotal := 50

	for _, value := range lines {

		dir := value[0]
		amountStr := value[1:]
		amount := str.ToInt(amountStr)

		if dir == 'L' {
			amount *= -1
		}

		turningTotal += amount

		if turningTotal%100 == 0 {
			res += 1
		}

	}

	return res
}

func solvePart2(lines []string) int {
	res := 0

	turningTotal := 50

	for _, value := range lines {

		dir := value[0]
		amountStr := value[1:]
		amount := str.ToInt(amountStr)

		for range amount {
			if dir == 'L' {
				turningTotal -= 1
			} else {
				turningTotal += 1
			}

			if turningTotal%100 == 0 {
				res += 1
			}

		}

	}

	return res
}
