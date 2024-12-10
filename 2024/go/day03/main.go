package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
)

var (
	REGEXSTRING2 string = `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`
	matches      [][]string
)

//go:embed input.txt
var input string

func main() {
	log.Println(solvePart1(input))
	log.Println(solvePart2Better(input))
}

func solvePart1(input string) int {
	res := 0
	re := regexp.MustCompile(REGEXSTRING2)
	matches = re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])

		res += n1 * n2
	}

	return res
}

func solvePart2Better(input string) int {
	res := 0
	multi := true
	for _, match := range matches {
		if match[0] == "do()" {
			multi = true
		} else if match[0] == "don't()" {
			multi = false
		} else {
			if multi {
				n1, _ := strconv.Atoi(match[1])
				n2, _ := strconv.Atoi(match[2])

				res += n1 * n2

			}
		}
	}

	return res
}

// func solvePart2(input string) int {
// 	res := 0
// 	input = strings.ReplaceAll(input, "\n", "")
//
// 	re := regexp.MustCompile(REGEXSTRING)
// 	indices := re.FindAllIndex([]byte(input), -1)
// 	matches := re.FindAllStringSubmatch(input, -1)
//
// 	muls := []string{}
//
// 	prev := 0
// 	for _, match := range indices {
// 		muls = append(muls, input[prev:match[1]])
// 		prev = match[1]
// 	}
//
// 	mutli := true
// 	for i, v := range muls {
// 		dont_index := strings.Index(v, "don't()")
// 		do_index := strings.Index(v, "do()")
//
// 		if dont_index != -1 && do_index == -1 {
// 			mutli = false
// 		} else if dont_index == -1 && do_index != -1 {
// 			mutli = true
// 		} else if dont_index != -1 && do_index != -1 {
// 			if do_index > dont_index {
// 				mutli = true
// 			} else {
// 				mutli = false
// 			}
// 		}
//
// 		if mutli {
// 			n1, err := strconv.Atoi(matches[i][1])
// 			if err != nil {
// 				log.Fatal(err)
// 			}
//
// 			n2, err := strconv.Atoi(matches[i][2])
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			res += n1 * n2
// 		}
//
// 	}
//
// 	return res
// }
