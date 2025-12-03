package main

import (
	"testing"
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

func TestPart1(t *testing.T) {
	input := []string{
		`987654321111111
811111111111119
234234234234278
818181911112111`,
	}
	exp := 357
	res := solvePart1(input)
	if res != exp {
		t.Errorf("solvePart1() = %d, exp %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"",
	}
	exp := 0
	res := solvePart2(input)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}
