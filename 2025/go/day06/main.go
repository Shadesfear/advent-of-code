package main

import (
	"log"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/str"
)

func main() {
	lines, err := files.ReadInputLines("../../inputs/day06.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

type Op func(a, b int) int

type Compuation struct {
	numbers []int
	op      Op
}

func (c Compuation) Compute() int {
	first := c.numbers[0]
	for i := 1; i < len(c.numbers); i++ {
		first = c.op(first, c.numbers[i])
	}
	return first
}

func Plus(a, b int) int {
	return a + b
}

func Mult(a, b int) int {
	return a * b
}

func Parse(lines []string) []Compuation {
	ops := strings.Fields(lines[len(lines)-1])
	nCols := len(ops)
	comps := make([]Compuation, nCols)

	fmap := map[string]Op{
		"+": Plus,
		"*": Mult,
	}

	for key, op := range ops {
		comps[key].op = fmap[op]
	}

	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		lineNumbers := strings.Fields(line)
		for j := range lineNumbers {
			lineNumber := str.ToInt(lineNumbers[j])
			comps[j].numbers = append(comps[j].numbers, lineNumber)
		}
	}

	return comps
}

func solvePart1(lines []string) int {
	res := 0

	comps := Parse(lines)

	for _, comp := range comps {
		res += comp.Compute()
	}

	return res
}

func solvePart2(lines []string) int {
	res := 0
	col := 0

	ops := strings.Fields(lines[len(lines)-1])
	nOps := len(ops)

	comps := make([]Compuation, nOps)
	curComp := 0
	lineLen := len(lines[0])
	nLines := len(lines) - 1

	for col < lineLen {
		n := 0
		empty := 0
		for r := range nLines {
			char := lines[r][col]
			if char == ' ' {
				empty++
				continue
			}
			digit := int(char) - int('0')
			if n == 0 {
				n = digit
			} else {
				n = n*10 + digit
			}
		}
		if empty == nLines {
			curComp++
			col++
			continue
		}
		comps[curComp].numbers = append(comps[curComp].numbers, n)
		col++
	}

	fmap := map[string]Op{
		"+": Plus,
		"*": Mult,
	}

	for key, op := range ops {
		comps[key].op = fmap[op]
	}

	for _, comp := range comps {
		res += comp.Compute()
	}

	return res
}
