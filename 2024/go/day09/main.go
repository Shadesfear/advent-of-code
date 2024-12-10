package main

import (
	"log"

	"github.com/shadesfear/aoc-lib-go/files"
	// "github.com/shadesfear/aoc-lib-go/str"
	// "github.com/shadesfear/aoc-lib-go/str"
)

type Pair struct {
	a any
	b any
}

func main() {
	lines, err := files.ReadInputFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

func Files(input string) []int {
	res := []int{}

	id := 0
	for i, c := range input {
		repeatCount := int(c - '0')
		if repeatCount < 0 {
			continue
		}
		if i%2 == 0 {
			for i := 0; i < repeatCount; i++ {
				res = append(res, id)
			}
			id++
		} else {
			for i := 0; i < repeatCount; i++ {
				res = append(res, -1)
			}
		}
	}
	return res
}

func solvePart1(lines string) int {
	res := 0
	f := Files(lines)

	left, right := 0, len(f)-1
	for left != right {

		if f[left] != -1 {
			left++
			continue
		}

		if f[right] == -1 {
			right--
			continue
		}

		f[left], f[right] = f[right], f[left]
	}

	for i, v := range f {
		if v == -1 {
			break
		}
		res += i * v
	}
	return res
}

func FindBlocks(ids []int) []Pair {
	blocks := []Pair{}
	inPair := false

	start := 0
	current := -2

	i := 0
	for i < len(ids) {
		if !inPair && ids[i] != -1 {
			inPair = true
			current = ids[i]
			start = i
		}
		if inPair && ids[i] == -1 {
			inPair = false
			blocks = append(blocks, Pair{start, i})
		} else if inPair && ids[i] != current {
			inPair = false
			blocks = append(blocks, Pair{start, i})
			current = -2
			continue

		}
		i++
	}
	blocks = append(blocks, Pair{start, i})
	return blocks
}

func FindFreeSpace(ids []int) []Pair {
	blocks := []Pair{}
	inFree := false

	start := 0

	for i := 0; i < len(ids); i++ {
		if !inFree && ids[i] == -1 {
			inFree = true
			start = i
		}
		if inFree && ids[i] != -1 {
			inFree = false
			blocks = append(blocks, Pair{start, i})
		}
	}
	return blocks
}

func solvePart2(lines string) int {
	res := 0
	f := Files(lines)
	blocks := FindBlocks(f)

	for i := len(blocks) - 1; i >= 0; i-- {

		block := blocks[i]
		free := FindFreeSpace(f)
		moved := false
		for _, fr := range free {
			if !(block.a.(int) > fr.a.(int) && block.b.(int) > fr.b.(int)) {
				continue
			}
			if fr.b.(int)-fr.a.(int) >= block.b.(int)-block.a.(int) {
				putIntoFreeSpace(f, fr, block)
				moved = true
				break
			}
		}
		if moved {
			continue
		}
	}

	for i, v := range f {
		if v == -1 {
			continue
		}
		res += i * v
	}

	return res
}

func putIntoFreeSpace(f []int, fr, block Pair) []int {
	dig := f[block.a.(int)]
	c := 0
	for i := fr.a.(int); i < fr.b.(int) && c < block.b.(int)-block.a.(int); i++ {
		f[i] = dig
		c++
	}

	for i := block.a.(int); i < block.b.(int); i++ {
		f[i] = -1
	}
	return f
}
