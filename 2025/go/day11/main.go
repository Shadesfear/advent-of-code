package main

import (
	"log"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"

	"github.com/shadesfear/aoc-lib-go/datastructures"
)

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

	graph := parse(lines)
	q := datastructures.NewQueue[string]()
	q.Enqueue("you")

	for !q.IsEmpty() {
		cur, _ := q.Dequeue()

		if cur == "out" {
			res++
		}

		for _, next := range graph[cur] {
			q.Enqueue(next)
		}

	}

	return res
}

type MemKey struct {
	node string
	dac  bool
	fft  bool
}

func recur(cur string, graph map[string][]string, mem map[MemKey]int, dac, fft bool) int {

	if cur == "out" && dac && fft {
		return 1
	}

	memkey := MemKey{cur, dac, fft}

	if val, ok := mem[memkey]; ok {
		return val
	}

	if cur == "dac" {
		dac = true
	}
	if cur == "fft" {
		fft = true
	}

	res := 0
	for _, next := range graph[cur] {
		res += recur(next, graph, mem, dac, fft)
	}
	mem[memkey] = res

	return res
}

func solvePart2(lines []string) int {
	res := 0

	graph := parse(lines)

	res = recur("svr", graph, map[MemKey]int{}, false, false)

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
