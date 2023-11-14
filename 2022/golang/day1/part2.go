package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func normal() {

	content, _ := ioutil.ReadFile("input.txt")
	str := string(content)
	str = strings.ReplaceAll(str, "\r", "")

	biggest := 0
	second := 0
	third := 0

	curr := 0

	lines := strings.Split(str, "\n")

	for _, line := range lines {
		snack, err := strconv.Atoi(line)
		curr += snack

		if err != nil {
			if curr > third {
				third = curr
			}
			if third > second {
				third, second = second, third
			}
			if second > biggest {
				second, biggest = biggest, second
			}

			curr = 0

		}

	}
	fmt.Println(biggest + second + third)
}

func lists() {
	content, _ := ioutil.ReadFile("input.txt")
	str := string(content)
	str = strings.ReplaceAll(str, "\r", "")

	chunks := strings.Split(str, "\n\n")
	num_chunks := []int{}
	result := 0

	for _, chunk := range chunks {
		chunkSum := 0
		nums := strings.Split(chunk, "\n")
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			chunkSum += n
		}
		num_chunks = append(num_chunks, chunkSum)
	}
	sort.Ints(num_chunks)

	for _, num := range num_chunks[len(num_chunks)-3:] {
		result += num
	}
	fmt.Println(result)
}

func main() {
	lists()
}
