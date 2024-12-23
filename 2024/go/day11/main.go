package main

import (
	"log"
	"math"

	"github.com/shadesfear/aoc-lib-go/files"
)

func main() {
	lines, err := files.ReadInputFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}

	input, err := files.ParseInputToInts(lines)
	if err != nil {
		log.Fatal(err)
	}
	part1 := solvePart1(input)
	log.Println(part1)

	part2 := solvePart2(input)
	log.Println(part2)
}

func NextNumber(i int) (int, int, bool) {
	if i == 0 {
		return 1, 0, false
	}

	nDigits := int(math.Log10(float64(i)) + 1)

	if nDigits%2 == 0 {
		n1 := i / int(math.Pow10(nDigits/2))
		n2 := i % int(math.Pow10(nDigits/2))

		return n1, n2, true
	}

	return i * 2024, 0, false
}

func solvePart1(numbers []int) int {
	next := []int{}
	for i := 0; i < 25; i++ {
		next = []int{}
		for _, number := range numbers {
			n1, n2, n2ok := NextNumber(number)
			if n2ok {
				next = append(next, n2)
			}
			next = append(next, n1)

		}
		numbers = next
	}

	return len(next)
}

type Pair struct {
	a, b int
}

var memory map[Pair]int = map[Pair]int{}

func Iter(x, n int) int {
	ns, ok := memory[Pair{x, n}]
	if ok {
		return ns
	}

	res := 0

	nDigits := int(math.Log10(float64(x)) + 1)
	if n == 0 {
		return 1
	}
	if x == 0 {
		res = Iter(1, n-1)
	} else if nDigits%2 == 0 {
		n1 := x / int(math.Pow10(nDigits/2))
		n2 := x % int(math.Pow10(nDigits/2))
		res += Iter(n1, n-1)
		res += Iter(n2, n-1)

	} else {
		res = Iter(2024*x, n-1)
	}
	memory[Pair{x, n}] = res
	return memory[Pair{x, n}]
}

func solvePart2(numbers []int) int {
	res := 0
	for _, num := range numbers {
		res += Iter(num, 10)
	}
	return res
}
