package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/str"
	// "github.com/shadesfear/aoc-lib-go/math"
)

func main() {
	lines, err := files.ReadInputFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

type Button struct {
	X int
	Y int
}

type ClawMachine struct {
	A    Button
	B    Button
	X, Y int
}

func (cm ClawMachine) Print() {
	log.Printf("Button A: X+%d, Y+%d", cm.A.X, cm.A.Y)
	log.Printf("Button B: X+%d, Y+%d", cm.B.X, cm.B.Y)
	log.Printf("Prize: X=%d, Y=%d", cm.X, cm.Y)
}

func parse(input string) []ClawMachine {
	ClawMachines := []ClawMachine{}
	cms := strings.Split(input, "\n\n")
	r, _ := regexp.Compile(`(\d+)`)

	for _, cmstring := range cms {
		if len(cmstring) <= 1 {
			continue
		}

		cm := strings.Split(cmstring, "\n")
		if len(cm) < 3 {
			log.Fatal("Should be 3 lines")
		}

		an := r.FindAllString(cm[0], -1)
		bn := r.FindAllString(cm[1], -1)
		pn := r.FindAllString(cm[2], -1)
		aButton := Button{(str.ToInt(an[0])), str.ToInt(an[1])}
		bButton := Button{str.ToInt(bn[0]), str.ToInt(bn[1])}
		clawMachine := ClawMachine{
			aButton,
			bButton,
			str.ToInt(pn[0]),
			str.ToInt(pn[1]),
		}
		ClawMachines = append(ClawMachines, clawMachine)

	}

	return ClawMachines
}

func SolveClawMachine(cm ClawMachine) int {
	cm.Print()
	x := cm.X
	y := cm.Y
	best := 5000

	for i := range 101 {
		for j := range 101 {
			nextX := cm.A.X*(i) + cm.B.X*(j)
			nextY := cm.A.Y*(i) + cm.B.Y*(j)
			if x == nextX && y == nextY {
				best = min(best, 3*i+j)
			}
		}
	}

	if best < 5000 {
		return best
	}

	return 0
}

var adds int = 10000000000000

func SolveClawMachine2(cm ClawMachine) int64 {
	Px := cm.X + adds
	Py := cm.Y + adds

	den := (cm.A.X*cm.B.Y - cm.B.X*cm.A.Y)

	aPress := (Px*cm.B.Y - cm.B.X*Py)
	bPress := (cm.A.X*Py - Px*cm.A.Y)

	if aPress%den == 0 && bPress%den == 0 {
		return 3*int64(aPress/den) + int64(bPress/den)
	}

	return 0
}

func solvePart1(lines string) int {
	res := 0
	cms := parse(lines)

	for _, cm := range cms {
		res += SolveClawMachine(cm)
	}

	return res
}

func solvePart2(lines string) int64 {
	var res int64

	cms := parse(lines)

	for _, cm := range cms {
		res += SolveClawMachine2(cm)
	}

	return res
}
