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

type Dir4 Point

func (d Dir4) Opposite() Dir4 {
	return Dir4{d.X * -1, d.Y * -1}
}

var (
	Up    Dir4 = Dir4(NewPoint(0, -1))
	Right Dir4 = Dir4(NewPoint(1, 0))
	Down  Dir4 = Dir4(NewPoint(0, 1))
	Left  Dir4 = Dir4(NewPoint(-1, 0))
)

func (d Dir4) RotateCW() Dir4 {
	if d == Up {
		return Right
	}

	if d == Right {
		return Down
	}

	if d == Down {
		return Left
	}
	return Up
}

func (d Dir4) RotateCCW() Dir4 {
	if d == Up {
		return Left
	}

	if d == Right {
		return Up
	}

	if d == Down {
		return Right
	}
	return Down
}

func (d Dir4) Point() Point {
	return Point(d)
}

var Dirs4 []Dir4 = []Dir4{
	Dir4(NewPoint(1, 0)),
	Dir4(NewPoint(0, 1)),
	Dir4(NewPoint(-1, 0)),
	Dir4(NewPoint(0, -1)),
}

func (p Point) MoveDir(d Dir4) Point {
	return p.Add(d.Point())
}

func (p Point) ManhattenDistance(other Point) int {
	return myMath.Abs(p.X-other.X) + myMath.Abs(p.Y-other.Y)
}

func (p Point) Det(other Point) int {
	return p.X*other.Y - p.Y*other.X
}

func Shoelace(coords []Point) float64 {
	sum := 0
	for i := 0; i < len(coords)-1; i++ {
		sum += coords[i].Det(coords[i+1])
	}
	sum += coords[len(coords)-1].Det(coords[0])
	return float64(sum) / 2
}

func (p Point) EuclideanDistance(other Point) float64 {
	return math.Sqrt(math.Pow(float64(p.X-other.X), 2) + math.Pow(float64(p.Y-other.Y), 2))
}

func (p Point) Move(dx, dy int) Point {
	return Point{p.X + dx, p.Y + dy}
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
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
