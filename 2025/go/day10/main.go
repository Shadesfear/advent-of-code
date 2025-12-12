package main

import (
	"log"
	m "math"
	"os/exec"
	"strings"

	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/math"
	"github.com/shadesfear/aoc-lib-go/str"

	// "github.com/shadesfear/aoc-lib-go/math"

	"github.com/shadesfear/aoc-lib-go/datastructures"
)

type Point = datastructures.Point

type Button = int

type Machine struct {
	State        int
	Diagram      int
	Width        int
	Buttons      []Button
	Joltage      []int
	JoltageState []int
}

func activeBits(n int) []int {
	var positions []int
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			positions = append(positions, i)
		}
		n >>= 1
	}
	return positions
}

func parseBinary(s string) int {
	var n int
	for _, c := range s {
		n <<= 1
		if c == '#' {
			n |= 1
		}
	}
	return n
}

func toggleBits(n int, width int, positions ...int) int {
	for _, p := range positions {
		n ^= 1 << (width - 1 - p)
	}
	return n
}

func parseIntList(s string) []int {
	var result []int
	for num := range strings.SplitSeq(s, ",") {
		result = append(result, str.ToInt(num))
	}
	return result
}

func stripDelimiters(s string) string {
	return s[1 : len(s)-1]
}

func NewMachine(line string) Machine {
	parts := strings.Split(line, " ")

	diagramString := stripDelimiters(parts[0])
	diagram := parseBinary(diagramString)
	width := len(diagramString)

	var buttons []Button
	for _, s := range parts[1 : len(parts)-1] {
		ns := parseIntList(stripDelimiters(s))
		button := toggleBits(0, width, ns...)
		buttons = append(buttons, button)

	}

	joltage := parseIntList(stripDelimiters(parts[len(parts)-1]))

	return Machine{Diagram: diagram, Width: width, Buttons: buttons, Joltage: joltage}
}

func (b Machine) PrintDiagram() {
	log.Printf("%0*b\n", b.Width, b.Diagram)
}

func (b Machine) PrintButton(n int) {
	log.Print(activeBits(b.Buttons[n]), " ")
}

func (b Machine) Print() {
	log.Printf("State:   %0*b", b.Width, b.State)
	log.Printf("Diagram: %0*b", b.Width, b.Diagram)
	log.Printf("Width:   %d", b.Width)
	log.Printf("Buttons:")
	for i, btn := range b.Buttons {
		log.Printf("  [%d]: bits %v", i, activeBits(btn))
	}
}

func (b *Machine) ClickButton(button int) {
	b.State ^= b.Buttons[button]
}

func (b *Machine) SolveBruteforce() int {
	n := math.Pow(2, len(b.Buttons))
	smallest := m.MaxInt
	for i := range n {
		active := activeBits(i)
		if len(active) > smallest {
			continue
		}
		for _, idx := range active {
			b.ClickButton(idx)
		}
		if b.State == b.Diagram {
			if len(active) < smallest {
				smallest = len(active)
			}
		}

		b.State = 0
	}
	return smallest
}

func (b *Machine) SolveJoltage() int {
	numButtons := len(b.Buttons)
	numCounters := len(b.Joltage)
	buttonAffects := make([][]int, numButtons)

	// Just parse this in newMachine, its the puzzle input goddammit
	for i, btn := range b.Buttons {
		for j := 0; j < numCounters; j++ {
			if btn&(1<<(b.Width-1-j)) != 0 {
				buttonAffects[i] = append(buttonAffects[i], j)
			}
		}
	}

	// log.Println(buttonAffects)

	maxTarget := 0
	for _, j := range b.Joltage {
		if j > maxTarget {
			maxTarget = j
		}
	}

	best := m.MaxInt
	counters := make([]int, numCounters)

	var solve func(buttonIdx int, totalPresses int)
	solve = func(buttonIdx int, totalPresses int) {
		if totalPresses >= best {
			return
		}

		for i, c := range counters {
			if c > b.Joltage[i] {
				return
			}
		}

		if buttonIdx == numButtons {
			for i, c := range counters {
				if c != b.Joltage[i] {
					return
				}
			}
			best = totalPresses
			return
		}

		maxPresses := maxTarget
		for _, counterIdx := range buttonAffects[buttonIdx] {
			remaining := b.Joltage[counterIdx] - counters[counterIdx]
			if remaining < maxPresses {
				maxPresses = remaining
			}
		}

		for presses := 0; presses <= maxPresses; presses++ {
			for _, counterIdx := range buttonAffects[buttonIdx] {
				counters[counterIdx] += presses
			}

			solve(buttonIdx+1, totalPresses+presses)

			for _, counterIdx := range buttonAffects[buttonIdx] {
				counters[counterIdx] -= presses
			}
		}
	}
	solve(0, 0)

	// log.Println(best)

	return best
}

func main() {
	lines, err := files.ReadInputLines("../../inputs/day10.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

func solvePart1(lines []string) int {
	log.SetFlags(0)
	res := 0
	for _, line := range lines {
		machine := NewMachine(line)
		// machine.Print()
		res += machine.SolveBruteforce()
	}

	return res
}

func solvePart2(lines []string) int {
	res := 0
	for i, line := range lines {
		if line == "" {
			continue
		}
		cmd := exec.Command("uv", "run", "solver.py", line)
		out, _ := cmd.Output()
		ans := str.ToInt(strings.TrimSpace(string(out)))
		log.Printf("LineNumber: %d, line: %s, out: %v,  ans: %d", i, line, out, ans)

		res += ans
	}
	return res
}
