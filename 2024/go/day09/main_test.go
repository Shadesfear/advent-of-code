package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "2333133121414131402\n"

	exp := 1928
	res := solvePart1(input)

	if exp != res {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func TestPart2(t *testing.T) {
	input := "2333133121414131402\n"

	exp := 2858
	res := solvePart2(input)

	if exp != res {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}
