package math

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		Name string
		A    int
		B    int
		Want int
	}{
		{"1", 1071, 462, 21},
		{"2", 12, 18, 6},
		{"3", 100, 60, 20},
		{"4", 8, 10, 2},
		{"5", 20, 28, 4},
		{"6", 98, 56, 14},
		{"6", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ans := GCD(tt.A, tt.B)
			if ans != tt.Want {
				t.Errorf("got %d, want %d", ans, tt.Want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		Name string
		A    int
		Want int
	}{
		{"1", 5, 5},
		{"2", -5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ans := Abs(tt.A)
			if ans != tt.Want {
				t.Errorf("got %d, want %d", ans, tt.Want)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		Name string
		A    int
		B    int
		Want int
	}{
		{"1", 21, 6, 42},
		{"2", 2, 3, 6},
		{"3", 6, 10, 30},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ans := LCM(tt.A, tt.B)
			if ans != tt.Want {
				t.Errorf("got %d, want %d", ans, tt.Want)
			}
		})
	}
}

func TestPerm(t *testing.T) {
	l := []int{1, 2, 3}
	perms := Permutations(l)
	fmt.Println(perms)
}
