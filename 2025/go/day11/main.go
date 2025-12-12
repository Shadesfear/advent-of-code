package main

import (
	"log"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"

	// "github.com/shadesfear/aoc-lib-go/math"

	"github.com/shadesfear/aoc-lib-go/datastructures"
)

type Point = datastructures.Point

func parse(lines []string) map[string][]string {
	tree := map[string][]string{}

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		cur := parts[0]
		children := strings.Split(parts[1], " ")
		tree[cur] = children
	}

	return tree

}

func solvePart1(lines []string) int {
	res := 0

	tree := parse(lines)

	q := datastructures.NewQueue[string]()
	q.Enqueue("you")

	for !q.IsEmpty() {
		cur, _ := q.Dequeue()

		if cur == "out" {
			res++
		}

		for _, next := range tree[cur] {
			q.Enqueue(next)
		}

	}

	return res
}

func solvePart2(lines []string) int {
	res := 0

	tree := parse(lines)

	s := datastructures.NewStack[string]()
	s.Push("svr")

	fft, dac := false, false
	for !s.IsEmpty() {
		cur, _ := s.Pop()
		if cur == "fft" {
			fft = true
		}
		if cur == "dac" {
			dac = true
		}

		if cur == "out" {
			if fft && dac {
				res++
			}
			fft, dac = false, false
		}

		for _, next := range tree[cur] {
			s.Push(next)
		}

	}

	return res
}

func main() {
	lines, err := files.ReadInputLines("../../inputs/day11.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}
