package main

import (
	"testing"

	"github.com/shadesfear/aoc-lib-go/str"
)

func TestPowerBankLargest(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			ans := PowerBankLargest(tt.in)
			if ans != tt.out {
				t.Errorf("got %v, want %v", ans, tt.out)
			}
		})
	}
}

func TestRecurse(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, tt := range tests {

		t.Run(tt.in, func(t *testing.T) {
			ans := recurse(tt.in, 12)
			if ans != tt.out {
				t.Errorf("got %v, want %v", ans, tt.out)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`

	lines := str.SplitLines(input)

	exp := 357
	res := solvePart1(lines)
	if res != exp {
		t.Errorf("solvePart1() = %d, exp %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`

	lines := str.SplitLines(input)
	exp := 3121910778619
	res := solvePart2(lines)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}
