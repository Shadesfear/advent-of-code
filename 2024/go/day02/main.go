package main

import (
	"log"

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

	part2C := solvePart2Clever(lines)
	log.Println(part2C)
}

func solvePart1(lines []string) int {
	res := 0

	for _, line := range lines {
		numbers, err := files.ParseInputToInts(line)
		if err != nil {
			log.Fatal(err)
		}
		safe, _ := lineSolver(numbers)
		if !safe {
			continue
		}

		res += 1

	}
	return res
}

func lineSolver(numbers []int) (bool, []int) {
	inc := true
	dec := true
	safe := true
	problems := []int{}
	for i := 0; i < len(numbers)-1; i++ {

		if numbers[i] < numbers[i+1] {
			dec = false
		} else if numbers[i] > numbers[i+1] {
			inc = false
		}

		diff := math.Abs(numbers[i+1] - numbers[i])
		if diff > 3 || diff < 1 {
			problems = append(problems, i)
			problems = append(problems, i+1)
			safe = false
		}
	}

	if !(inc || dec) {
		safe = false
		problems = append(problems, 0)
	}

	return safe, problems
}

func generate(numbers []int) [][]int {
	res := [][]int{}
	for i := 0; i < len(numbers); i++ {
		local := []int{}

		for j, number := range numbers {
			if i != j {
				local = append(local, number)
			}
		}
		res = append(res, local)
	}
	return res
}

func removeIdx(numbers []int, idx int) []int {
	return append(numbers[:idx], numbers[idx+1:]...)
}

func solvePart2(lines []string) int {
	res := 0

	for _, line := range lines {
		numbers, err := files.ParseInputToInts(line)
		if err != nil {
			log.Fatal(err)
		}

		safe, _ := lineSolver(numbers)
		if !safe {
			subLines := generate(numbers)
			ok := false

			for _, subLine := range subLines {
				subSafe, _ := lineSolver(subLine)
				if subSafe {
					ok = true
					break
				}
			}

			if !ok {
				continue
			}
		}
		res += 1
	}
	return res
}

func solvePart2Clever(lines []string) int {
	res := 0

	for _, line := range lines {
		numbers, err := files.ParseInputToInts(line)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(line)

		safe, problems := lineSolver(numbers)
		log.Printf("   %v - %v\n", safe, problems)

		if !safe {
			ok := false

			for _, idx := range problems {
				subLine := removeIdx(numbers, idx)
				subSafe, _ := lineSolver(subLine)
				if subSafe {
					ok = true
					break
				}
			}

			if !ok {
				continue
			}
		}
		res += 1
	}
	return res
}
