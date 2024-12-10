package main

import (
	"log"
	"slices"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/str"
)

func main() {
	input, err := files.ReadInputFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}

	part1 := solvePart1(input)
	log.Println(part1)

	part2 := solvePart2(input)
	log.Println(part2)
}

func parse(input string) (map[int][]int, [][]int) {
	RULES := map[int][]int{}
	UPDATES := [][]int{}

	splits := strings.Split(input, "\n\n")
	rules := str.SplitLines(splits[0])
	updates := str.SplitLines(splits[1])

	for _, rule := range rules {
		numbers := strings.Split(rule, "|")
		n1 := str.ToInt(numbers[0])
		n2 := str.ToInt(numbers[1])

		v, ok := RULES[n1]
		if !ok {
			RULES[n1] = []int{}
		}
		RULES[n1] = append(v, n2)

	}

	for _, update := range updates {
		numStrings := strings.Split(update, ",")
		nums := []int{}
		for _, ns := range numStrings {
			if len(ns) < 1 {
				continue
			}
			num := str.ToInt(ns)
			nums = append(nums, num)
		}
		if len(nums) > 0 {
			UPDATES = append(UPDATES, nums)
		}

	}

	return RULES, UPDATES
}

func correct(update []int, rules map[int][]int) bool {
	for i := 0; i < len(update)-1; i++ {
		candidate := update[i]
		candidateRule := rules[candidate]
		for k := i + 1; k < len(update); k++ {
			if !slices.Contains(candidateRule, update[k]) {
				return false
			}
		}

	}
	return true
}

func solvePart1(input string) int {
	res := 0
	rules, updates := parse(input)
	for _, update := range updates {
		isCorrect := correct(update, rules)
		if isCorrect {
			res += update[len(update)/2]
		}
	}
	return res
}

func solvePart2(input string) int {
	res := 0
	rules, updates := parse(input)
	for _, update := range updates {

		if correct(update, rules) {
			continue
		}

		slices.SortFunc(update, func(a, b int) int {
			rule := rules[a]
			if slices.Contains(rule, b) {
				return -1
			}
			return 1
		})

		res += update[len(update)/2]
	}
	return res
}
