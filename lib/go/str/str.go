package str

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func SplitLines(input string) []string {
	return strings.Split(input, "\n")
}

func RemoveEmpty(strings []string) []string {
	out := []string{}
	for _, str := range strings {
		if len(str) > 0 {
			out = append(out, str)
		}
	}
	return out
}

func ToInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		log.Printf("Got error with input %s", input)
		panic(err)
	}
	return res
}

func ToInt64(input string) int64 {
	res, err := strconv.Atoi(input)
	if err != nil {
		log.Printf("Got error with input %s", input)
		panic(err)
	}
	return int64(res)
}

func PrettyPrintGrid[T any](grid [][]T) {
	for _, row := range grid {
		for _, cell := range row {
			if r, ok := any(cell).(rune); ok {
				fmt.Printf("%c", r)
			} else {
				fmt.Printf("%v", cell)
			}
		}
		fmt.Println()
	}
}

// func PrettyPrintGrid(grid [][]rune) {
// 	for _, row := range grid {
// 		for _, cell := range row {
// 			fmt.Printf("%c ", cell) // Print each rune with a space
// 		}
// 		fmt.Println() // Move to the next line after each row
// 	}
// }
