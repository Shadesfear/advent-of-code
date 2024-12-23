package main

import (
	"log"
	"slices"
	"testing"

	"github.com/shadesfear/aoc-lib-go/datastructures"
	"github.com/shadesfear/aoc-lib-go/files"
)

func TestPart1Single(t *testing.T) {
	inputs := []struct {
		input string
		exp   int
	}{
		{
			`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			772,
		},
	}

	for _, tt := range inputs {
		t.Run(tt.input, func(t *testing.T) {
			inp := files.StringToLines(tt.input)
			s := solvePart1(inp)
			if s != tt.exp {
				t.Errorf("got %d, want %d", s, tt.exp)
			}
		})
	}
}

func TestPart1Big(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	lines := files.StringToLines(input)
	res := solvePart1(lines)
	log.Println(res)
}

func TestAllInterior(t *testing.T) {
	input := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	lines := files.StringToLines(input)
	visited := []datastructures.Point{}

	blocks := [][]datastructures.Point{}

	for y, row := range lines {
		for x := range row {
			here := datastructures.NewPoint(x, y)
			if slices.Contains(visited, here) {
				continue
			}
			farm := bfs(lines, here)
			blocks = append(blocks, farm)
			visited = append(visited, farm...)

		}
	}

	res := allInterior(lines, blocks[0])

	log.Println(res)
}

func TestIsInte(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	lines := files.StringToLines(input)
	visited := []datastructures.Point{}

	blocks := [][]datastructures.Point{}

	for y, row := range lines {
		for x := range row {
			here := datastructures.NewPoint(x, y)
			if slices.Contains(visited, here) {
				continue
			}
			farm := bfs(lines, here)
			blocks = append(blocks, farm)
			visited = append(visited, farm...)

		}
	}

	res := isInteriorPoint(lines, blocks[1], datastructures.NewPoint(1, 1))
	log.Println(res)
}

func TestIsBoundary(t *testing.T) {
	input := `AAAA
AAAA
AAAA
EEEC`
	lines := files.StringToLines(input)
	res := isBoundary("A"[0], datastructures.NewPoint(1, 1), lines)
	log.Println(res)
}

// func TestPart1Single(t *testing.T) {
// 	inputs := []struct {
// 		input string
// 		exp   int
// 	}{
// 		{
// 			`OOOOO
// OXOXO
// OOOOO
// OXOXO
// OOOOO`,
// 			772,
// 		},
// 	}
//
// 	for _, tt := range inputs {
// 		t.Run(tt.input, func(t *testing.T) {
// 			inp := files.StringToLines(tt.input)
// 			s := solvePart1(inp)
// 			if s != tt.exp {
// 				t.Errorf("got %d, want %d", s, tt.exp)
// 			}
// 		})
// 	}
// }

func TestBfs(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	lines := files.StringToLines(input)
	start := datastructures.NewPoint(0, 0)
	res := bfs(lines, start)
	wants := []datastructures.Point{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 3, Y: 0},
	}
	if len(res) != 4 {
		t.Errorf("TestBfs(), returned nr of elements are wrong, got %d, exp %d", len(res), 4)
	}

	if !slices.Equal(wants, res) {
		t.Errorf("TestBfs(), does not have same elements, got %s, want %s", res, wants)
	}
}

func TestPart1(t *testing.T) {
	inputs := []struct {
		input string
		exp   int
	}{
		{
			`AAAA
BBCD
BBCC
EEEC`,
			140,
		},
		{
			`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			772,
		},
		{
			`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			1930,
		},
	}

	for _, tt := range inputs {
		t.Run(tt.input, func(t *testing.T) {
			inp := files.StringToLines(tt.input)
			s := solvePart1(inp)
			if s != tt.exp {
				t.Errorf("got %d, want %d", s, tt.exp)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	input := ``
	exp := 0
	lines := files.StringToLines(input)
	res := solvePart2(lines)
	if res != exp {
		t.Errorf("solvePart2() = %d, exp %d", res, exp)
	}
}
