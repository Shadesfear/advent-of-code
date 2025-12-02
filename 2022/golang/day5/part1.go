package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2
var test bool = true
var SPACING int = 4

func printStacks(stacks [][]string) {
	for _, stack := range stacks {
		fmt.Println(stack)
	}
}

func pop(stack []string) ([]string, string) {
	i := len(stack) - 1
	elem := stack[i]
	stack = stack[:i]
	return stack, elem
}

func main() {

	cols := 9
	filename := "../../inputs/day05.txt"

	if test {
		cols = 3
		filename = "test.txt"
	}
	content, _ := ioutil.ReadFile(filename)

	str := string(content)
	str = strings.ReplaceAll(str, "\r", "")

	inp := strings.Split(str, "\n\n")
	crates := inp[0]
	crates = strings.ReplaceAll(crates, "[", " ")
	crates = strings.ReplaceAll(crates, "]", " ")

	crates_rows := strings.Split(crates, "\n")
	stacks := make([][]string, cols)

	commands := strings.Split(inp[1], "\n")

	for i := 0; i < cols; i++ {
		idx := i*SPACING + 1
		stacks[i] = []string{}
		for j := 0; j < len(crates_rows)-1; j++ {
			row := crates_rows[j]
			if len(row) <= idx {
				continue
			}
			if string(row[idx]) == " " {
				continue
			}
			stacks[i] = append(stacks[i], string(row[idx]))
		}
	}

	for _, command := range commands {
		fmt.Println(command)
	}

	// fmt.Println(stacks)

	// printStacks(stacks)

	// _, _ = pop(stacks[0])

	// // stacks[0] = stack

	// printStacks(stacks)

	// fmt.Println(crates)

	// commands := strings.Split(inp[1], "\n")

	// for _, baj := range commands {
	// 	fmt.Println(baj)
	// }

}
