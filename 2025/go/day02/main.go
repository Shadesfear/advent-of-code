package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/str"

	// "github.com/shadesfear/aoc-lib-go/math"

	"github.com/shadesfear/aoc-lib-go/datastructures"
)

type Point = datastructures.Point

func init() {
	log.SetOutput(io.Discard)
}

func main() {
	// lines, err := files.ReadInputLines("../../inputs/day01.txt")
	log.SetOutput(os.Stdin)
	lines, err := files.ReadDayInput(2) // ("../../inputs/day01.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

type Range struct {
	left  int
	right int
}

func rangeFromString(input string) Range {
	input = strings.Replace(input, "\n", "", 1)
	spl := strings.Split(input, "-")
	left := str.ToInt(spl[0])
	right := str.ToInt(spl[1])
	r := Range{
		left,
		right,
	}
	return r
}

func rangesInvalidP1(r Range) []int {
	res := []int{}
	for i := r.left; i < r.right+1; i++ {
		str := strconv.Itoa(i)
		strLen := len(str)
		if strLen%2 != 0 {
			continue
		}

		middle := strLen / 2

		if str[:middle] == str[middle:] {
			res = append(res, i)
		}

	}
	return res
}

func invalidIDP2(n int) bool {
	strN := strconv.Itoa(n)
	strNLen := len(strN)
	for i := 2; i <= strNLen; i++ {
		if strNLen%i != 0 {
			continue
		}
		chunkSize := strNLen / i

		same := true
		for j := 0; j < strNLen-chunkSize; j += chunkSize {
			if strN[j:j+chunkSize] != strN[j+chunkSize:j+chunkSize*2] {
				same = false
			}

		}

		if same {
			return true
		}

	}

	return false

}

func rangesInvalidP2(r Range) []int {
	res := []int{}
	for i := r.left; i < r.right+1; i++ {
		if invalidIDP2(i) {
			res = append(res, i)
		}

	}
	return res
}

func solvePart1(line string) int {
	ranges := strings.Split(line, ",")
	res := 0

	for _, srange := range ranges {

		r := rangeFromString(srange)

		validIds := rangesInvalidP1(r)
		for _, validID := range validIds {
			res += validID
		}

	}

	return res
}

func solvePart2(line string) int {
	ranges := strings.Split(line, ",")
	res := 0

	for _, srange := range ranges {

		r := rangeFromString(srange)

		validIds := rangesInvalidP2(r)
		for _, validID := range validIds {
			res += validID
		}

	}

	return res
}
