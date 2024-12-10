package main

import (
	"log"
	"math"
	// "os"
	// "os/exec"
	// "time"

	"github.com/shadesfear/aoc-lib-go/datastructures"
	"github.com/shadesfear/aoc-lib-go/files"
	// "github.com/shadesfear/aoc-lib-go/math"
)

var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

var colors []string = []string{
	Red, Green, Yellow, Blue, Magenta, Cyan, Gray, White,
}

func main() {
	lines, err := files.ReadInputLines("input.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}
	part1 := solvePart1(lines)
	log.Println(part1)

	part2 := solvePart2(lines)
	log.Println(part2)
}

type Vector struct {
	X float64
	Y float64
}

func NewVector(x, y float64) Vector {
	return Vector{x, y}
}

func (p Vector) Norm() float64 {
	return math.Sqrt(math.Pow(float64(p.X), 2) + math.Pow(float64(p.Y), 2))
}

func (p Vector) Direction(other Vector) Vector {
	dir := NewVector(p.X-other.X, p.Y-other.Y)
	mag := dir.Norm()
	return NewVector(dir.X/mag, dir.Y/mag)
}

func (v Vector) EuclideanDistance(other Vector) float64 {
	return math.Sqrt(math.Pow(float64(v.X-other.X), 2) + math.Pow(float64(v.Y-other.Y), 2))
}

func (v Vector) MoveByDirection(amount float64, direction Vector) Vector {
	dx := direction.X * amount
	dy := direction.Y * amount
	return NewVector(v.X+dx, v.Y+dy)
}

func (v Vector) InBounds(x, y int) bool {
	return int(v.X) >= 0 && int(v.X) < x && int(v.Y) >= 0 && int(v.Y) < y
}

func FromPoint(p datastructures.Point) Vector {
	return NewVector(float64(p.X), float64(p.Y))
}

func (v Vector) ToPoint() datastructures.Point {
	X := int(math.Round(v.X))
	Y := int(math.Round(v.Y))
	return datastructures.NewPoint(X, Y)
}

func contains(list []datastructures.Point, target datastructures.Point) bool {
	for _, l := range list {
		if l.Equal(target) {
			return true
		}
	}
	return false
}

func solvePart1(lines []string) int {
	found := map[byte][]datastructures.Point{}
	rows, cols := len(lines), len(lines[0])

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if lines[y][x] != '.' {
				found[lines[y][x]] = append(found[lines[y][x]], datastructures.NewPoint(x, y))
			}
		}
	}

	antinodes := []datastructures.Point{}

	for _, f := range found {
		for i := 0; i < len(f); i++ {
			for j := i + 1; j < len(f); j++ {
				p1 := FromPoint(f[i])
				p2 := FromPoint(f[j])

				dis := p1.EuclideanDistance(p2)
				dir := p1.Direction(p2)
				dir2 := p2.Direction(p1)

				antinode := p1.MoveByDirection(dis, dir).ToPoint()
				antinode2 := p2.MoveByDirection(dis, dir2).ToPoint()

				if antinode.InBounds(rows, cols) && !contains(antinodes, antinode) {
					antinodes = append(antinodes, antinode)
				}

				if antinode2.InBounds(rows, cols) && !contains(antinodes, antinode2) {
					antinodes = append(antinodes, antinode2)
				}

			}
		}
	}

	return len(antinodes)
}

func solvePart2(lines []string) int {
	found := map[byte][]datastructures.Point{}
	rows, cols := len(lines), len(lines[0])

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if lines[y][x] != '.' {
				found[lines[y][x]] = append(found[lines[y][x]], datastructures.NewPoint(x, y))
			}
		}
	}

	antinodes := []datastructures.Point{}

	for _, f := range found {
		for i := 0; i < len(f); i++ {
			for j := i + 1; j < len(f); j++ {
				p1 := FromPoint(f[i])
				p2 := FromPoint(f[j])

				if !contains(antinodes, f[i]) && f[i].InBounds(rows, cols) {
					antinodes = append(antinodes, f[i])
				}

				if !contains(antinodes, f[j]) && f[j].InBounds(rows, cols) {
					antinodes = append(antinodes, f[j])
				}

				dis := p1.EuclideanDistance(p2)
				dir := p1.Direction(p2)

				for {

					p1 = p1.MoveByDirection(dis, dir)
					if !p1.ToPoint().InBounds(rows, cols) {
						break
					}

					if contains(antinodes, p1.ToPoint()) {
						continue
					}

					antinodes = append(antinodes, p1.ToPoint())

				}

				dir2 := p2.Direction(p1)

				for {

					p2 = p2.MoveByDirection(dis, dir2)
					if !p2.ToPoint().InBounds(rows, cols) {
						break
					}

					if contains(antinodes, p2.ToPoint()) {
						continue
					}

					antinodes = append(antinodes, p2.ToPoint())

				}

			}
		}
	}

	// for _, a := range antinodes {
	// 	// log.Println(a)
	// 	lines[a.Y] = lines[a.Y][:a.X] + "#" + lines[a.Y][a.X+1:]
	//
	// 	cmd := exec.Command("clear")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	//
	// 	for _, line := range lines {
	// 		log.Println(line)
	// 	}
	//
	// 	time.Sleep(time.Millisecond * 100)
	// }

	return len(antinodes)
}
