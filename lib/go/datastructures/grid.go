package datastructures

import (
	"fmt"
	"strings"
)

type Grid[T comparable] struct {
	Data [][]T
	Rows int
	Cols int
}

func NewGrid[T comparable](data [][]T) *Grid[T] {
	rows := len(data)
	cols := 0
	if rows > 0 {
		cols = len(data[0])
	}
	return &Grid[T]{Data: data, Rows: rows, Cols: cols}
}

func (g *Grid[T]) Get(p Point) T {
	if !p.InBounds(g.Rows, g.Cols) {
		panic("point out of bounds")
	}
	return g.Data[p.Y][p.X]
}

func (g *Grid[T]) Set(p Point, val T) {
	if !p.InBounds(g.Rows, g.Cols) {
		panic("point out of bounds")
	}
	g.Data[p.Y][p.X] = val
}

func (g *Grid[T]) InBounds(p Point) bool {
	return p.InBounds(g.Rows, g.Cols)
}

func (g *Grid[T]) Neighbours(p Point) []Point {
	var result []Point
	for _, n := range p.Neighbours() {
		if g.InBounds(n) {
			result = append(result, n)
		}
	}
	return result
}

func (g *Grid[T]) Neighbours8(p Point) []Point {
	var result []Point
	for _, n := range p.Neighbors8() {
		if g.InBounds(n) {
			result = append(result, n)
		}
	}
	return result
}

func (g *Grid[T]) Find(val T) (Point, bool) {
	for y := 0; y < g.Rows; y++ {
		for x := 0; x < g.Cols; x++ {
			if g.Data[y][x] == val {
				return Point{X: x, Y: y}, true
			}
		}
	}
	return Point{}, false
}

func (g *Grid[T]) FindAll(val T) []Point {
	var result []Point
	for y := 0; y < g.Rows; y++ {
		for x := 0; x < g.Cols; x++ {
			if g.Data[y][x] == val {
				result = append(result, Point{X: x, Y: y})
			}
		}
	}
	return result
}

func (g *Grid[T]) ForEach(fn func(Point, T)) {
	for y := 0; y < g.Rows; y++ {
		for x := 0; x < g.Cols; x++ {
			fn(Point{X: x, Y: y}, g.Data[y][x])
		}
	}
}

func (g *Grid[T]) Clone() *Grid[T] {
	data := make([][]T, g.Rows)
	for y := 0; y < g.Rows; y++ {
		data[y] = make([]T, g.Cols)
		copy(data[y], g.Data[y])
	}
	return NewGrid(data)
}

func (g *Grid[T]) String() string {
	var buf strings.Builder
	for y := 0; y < g.Rows; y++ {
		for x := 0; x < g.Cols; x++ {
			cell := g.Data[y][x]
			switch v := any(cell).(type) {
			case rune:
				buf.WriteRune(v)
			case byte:
				buf.WriteByte(v)
			default:
				fmt.Fprintf(&buf, "%v", cell)
			}
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (g *Grid[T]) Print() {
	fmt.Print(g.String())
}
