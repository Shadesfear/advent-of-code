package main

import (
	"log"
	"strings"
	// "sync"

	"github.com/shadesfear/aoc-lib-go/files"
	// "github.com/shadesfear/aoc-lib-go/math"
	// "github.com/shadesfear/aoc-lib-go/math"
	"github.com/shadesfear/aoc-lib-go/str"
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
}

func generateConfigurations(positions int, candidates []string) []string {
	var configs []string

	// Helper function to generate each configuration
	var generate func(current string)
	generate = func(current string) {
		// If we've filled all positions, add this configuration
		if len(current) == positions {
			configs = append(configs, current)
			return
		}

		// Try adding both 'a' and 'b' to the current position
		for _, c := range candidates {
			generate(current + string(c))
		}
		// generate(current + "+")
		// generate(current + "*")
	}

	// Start the generation with an empty string
	generate("")
	return configs
}

func Sum(nums []int) int {
	r := 0
	for _, n := range nums {
		r += n
	}
	return r
}

func solvePart1(lines []string) int {
	res := 0
	for _, line := range lines {
		split := strings.Split(line, ":")
		targetStr, numbersStr := split[0], split[1]
		target := str.ToInt(targetStr)
		numbers, err := files.ParseInputToInts(numbersStr)
		if err != nil {
			log.Fatal(err)
		}
		combinations := generateConfigurations(len(numbers)-1, []string{"*", "+"})

		for _, com := range combinations {
			result := numbers[0]
			for i := 1; i < len(numbers); i++ {
				if string(com[i-1]) == "*" {
					result = result * numbers[i]
				} else {
					result = result + numbers[i]
				}
			}
			if result == target {
				res += target
				break
			}
		}

	}
	return res
}

func concatenate(x, y int) int {
	pow := 10
	for y >= pow {
		pow *= 10
	}
	return x*pow + y
}

func solvePart2(lines []string) int {
	res := 0
	// var wg sync.WaitGroup
	for _, line := range lines {
		// wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		split := strings.Split(line, ":")
		targetStr, numbersStr := split[0], split[1]
		target := str.ToInt(targetStr)
		numbers, err := files.ParseInputToInts(numbersStr)
		if err != nil {
			log.Fatal(err)
		}

		aboveTarget := map[string]bool{}
		combinations := generateConfigurations(len(numbers)-1, []string{"*", "+", "."})

		for _, com := range combinations {
			result := numbers[0]
			breakEarly := false
			for i := 1; i < len(com); i++ {
				at := aboveTarget[com[:i]]
				if at {
					breakEarly = true
				}
			}
			if breakEarly {
				continue
			}
			for i := 1; i < len(numbers); i++ {

				if result > target {
					aboveTarget[com[:i]] = true
					break
				}

				if string(com[i-1]) == "*" {
					result = result * numbers[i]
				} else if string(com[i-1]) == "." {
					result = concatenate(result, numbers[i])
				} else {
					result = result + numbers[i]
				}

			}

			if result == target {
				res += target
				break
			}
		}
		// }()
	}

	// wg.Wait()
	return res
}
