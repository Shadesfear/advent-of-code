package datastructures

import (
	"fmt"
	"math"

	myMath "github.com/shadesfear/aoc-lib-go/math"
)

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (p Point) ManhattenDistance(other Point) int {
	return myMath.Abs(p.X-other.X) + myMath.Abs(p.Y-other.Y)
}

func (p Point) EuclideanDistance(other Point) float64 {
	return math.Sqrt(math.Pow(float64(p.X-other.X), 2) + math.Pow(float64(p.Y-other.Y), 2))
}

func (p Point) Move(dx, dy int) Point {
	return Point{p.X + dx, p.Y + dy}
}

func (p Point) Equal(other Point) bool {
	if p.X == other.X && p.Y == other.Y {
		return true
	}
	return false
}

func (p Point) Neighbours() []Point {
	return []Point{
		{X: p.X, Y: p.Y - 1}, // Up
		{X: p.X, Y: p.Y + 1}, // Down
		{X: p.X - 1, Y: p.Y}, // Left
		{X: p.X + 1, Y: p.Y}, // Right
	}
}

// Get 8 adjacent points (including diagonals)
func (p Point) Neighbors8() []Point {
	offsets := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	neighbors := make([]Point, 0, 8)
	for _, offset := range offsets {
		neighbors = append(neighbors, Point{X: p.X + offset[0], Y: p.Y + offset[1]})
	}

	return neighbors
}

func (p Point) InBounds(rows, cols int) bool {
	return p.X >= 0 && p.X < cols && p.Y >= 0 && p.Y < rows
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
