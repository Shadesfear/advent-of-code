package main

import (
	"strings"
	"testing"
)

func TestComputation(t *testing.T) {
	c := Compuation{
		[]int{123, 45, 6},
		func(a, b int) int {
			return a * b
		},
	}

	res := c.Compute()

	exp := 33210

	if res != exp {
		t.Errorf("TestComputation = %d, exp %d", res, exp)
	}

	c = Compuation{
		[]int{328, 64, 98},
		func(a, b int) int {
			return a + b
		},
	}

	res = c.Compute()

	exp = 490

	if res != exp {
		t.Errorf("TestComputation = %d, exp %d", res, exp)
	}

}

func TestPart1(t *testing.T) {
	input := `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +`
	lines := strings.Split(input, "\n")
	exp := 4277556
	res := solvePart1(lines)
	if res != exp {
		t.Errorf("solvePart1() = %d, exp %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}
	exp := 3263827
	res := solvePart2(input)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}
