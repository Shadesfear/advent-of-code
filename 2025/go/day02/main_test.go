package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"testing"
)

func init() {
	log.SetOutput(os.Stdout)

}

func TestRangeFromString(t *testing.T) {
	tests := []struct {
		in  string
		out Range
	}{
		{"11-22", Range{11, 22}},
		{"95-115", Range{95, 115}},
		{"998-1012", Range{998, 1012}},
		{"1188511880-1188511890", Range{1188511880, 1188511890}},
		{"222220-222224", Range{222220, 222224}},
		{"1698522-1698528", Range{1698522, 1698528}},
		{"446443-446449", Range{446443, 446449}},
		{"38593856-38593862", Range{38593856, 38593862}},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			ans := rangeFromString(tt.in)
			if ans.left != tt.out.left || ans.right != tt.out.right {
				t.Errorf("got %v, want %v", ans, tt.out)
			}
		})
	}

}

func TestInvalidP2(t *testing.T) {

	tests := []struct {
		in  int
		out bool
	}{
		{11, true},
		{101, false},
		{115, false},
		{1001, false},
		{1012, false},
		{222220, false},
		{22, true},
		{99, true},
		{111, true},
		{999, true},
		{1010, true},
		{1188511885, true},
		{222222, true},
		{1698522, false},
		{446446, true},
		{38593859, true},
		{824824824, true},
		{2121212121, true},
		{222222, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			ans := invalidIDP2(tt.in)
			if tt.out != ans {
				t.Errorf("got %v, want %v", ans, tt.out)
			}
		})
	}

}

func TestRangesInvalidP1(t *testing.T) {
	tests := []struct {
		in  Range
		out []int
	}{

		{Range{11, 22}, []int{11, 22}},
		{Range{95, 115}, []int{99}},
		{Range{998, 1012}, []int{1010}},
		{Range{1188511880, 1188511890}, []int{1188511885}},
		{Range{222220, 222224}, []int{222222}},
		{Range{1698522, 1698528}, []int{}},
		{Range{446443, 446449}, []int{446446}},
		{Range{38593856, 38593862}, []int{38593859}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			ans := rangesInvalidP1(tt.in)
			if !slices.Equal(tt.out, ans) {
				t.Errorf("got %v, want %v", ans, tt.out)
			}
		})
	}
}

func TestRangesInvalidP2(t *testing.T) {
	tests := []struct {
		in  Range
		out []int
	}{

		{Range{11, 22}, []int{11, 22}},
		{Range{95, 115}, []int{99, 111}},
		{Range{998, 1012}, []int{999, 1010}},
		{Range{1188511880, 1188511890}, []int{1188511885}},
		{Range{222220, 222224}, []int{222222}},
		{Range{1698522, 1698528}, []int{}},
		{Range{446443, 446449}, []int{446446}},
		{Range{38593856, 38593862}, []int{38593859}},
		{Range{565653, 565659}, []int{565656}},
		{Range{824824821, 824824827}, []int{824824824}},
		{Range{2121212118, 2121212124}, []int{2121212121}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			ans := rangesInvalidP2(tt.in)
			if !slices.Equal(tt.out, ans) {
				t.Errorf("got %v, want %v", ans, tt.out)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	exp := 1227775554
	res := solvePart1(input)
	if res != exp {
		t.Errorf("solvePart1() = %d, exp %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	exp := 4174379265
	res := solvePart2(input)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}
