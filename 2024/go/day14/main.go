package main

import (
	"fmt"
	"log"
	"regexp"

	ds "github.com/shadesfear/aoc-lib-go/datastructures"
	"github.com/shadesfear/aoc-lib-go/files"
	"github.com/shadesfear/aoc-lib-go/str"
	// "github.com/shadesfear/aoc-lib-go/math"
)

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

var (
	width  int = 101
	height int = 103
	// At coordinate, just save the velocity
	points map[ds.Point][]ds.Point = map[ds.Point][]ds.Point{}
)

func parse(lines []string) {
	for _, line := range lines {
		r, _ := regexp.Compile(`(-?\d+)`)
		ns := r.FindAllString(line, -1)
		pos := ds.NewPoint(str.ToInt(ns[0]), str.ToInt(ns[1]))
		vel := ds.NewPoint(str.ToInt(ns[2]), str.ToInt(ns[3]))

		_, ok := points[pos]
		if !ok {
			points[pos] = []ds.Point{}
		}

		points[pos] = append(points[pos], vel)
	}
}

func printPoints() {
	for y := range height {
		fmt.Println()
		for x := range width {
			point := ds.NewPoint(x, y)
			_, ok := points[point]
			if ok {
				fmt.Printf("%d", len(points[point]))
				continue
			}
			fmt.Printf(".")

		}
	}
}

func move() {
	newPoints := map[ds.Point][]ds.Point{}
	for cur, vels := range points {
		for _, vel := range vels {
			moved := cur.Move(vel.X, vel.Y)

			newX := (moved.X + width) % width
			newY := (moved.Y + height) % height

			newLoc := ds.NewPoint(newX, newY)

			_, ok := newPoints[newLoc]
			if !ok {
				newPoints[newLoc] = []ds.Point{}
			}
			newPoints[newLoc] = append(newPoints[newLoc], vel)

		}
	}
	points = newPoints
}

func multiplyPositive(nums ...int) int {
	result := 1
	for _, num := range nums {
		if num > 0 {
			result *= num
		}
	}
	return result
}

func countInQuadrant() int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for point, vels := range points {
		if point.X == width/2 || point.Y == height/2 {
			continue
		}
		if point.X < width/2 && point.Y < height/2 {
			q1 += len(vels)
		}
		if point.X > width/2 && point.Y < height/2 {
			q2 += len(vels)
		}

		if point.X < width/2 && point.Y > height/2 {
			q3 += len(vels)
		}

		if point.X > width/2 && point.Y > height/2 {
			q4 += len(vels)
		}

	}

	return multiplyPositive(q1, q2, q3, q4)
}

func density(sX, sY, eX, eY int) float64 {
	area := (eX - sX) * (eY - sY)
	count := 0

	for point, vels := range points {
		if point.X > sX && point.X < eX && point.Y < eY && point.Y > sY {
			count += len(vels)
		}
	}
	return float64(count) / float64(area)
}

func solvePart1(lines []string) int {
	parse(lines)
	for range 100 {
		move()
	}
	return countInQuadrant()
}

func solvePart2(lines []string) int {
	parse(lines)
	for i := 0; i < 100000; i++ {

		seen := ds.NewSet[ds.Point]()
		move()
		for k := range points {
			seen.Add(k)
		}
		if seen.Len() == len(lines) {
			return i
		}

		// smallDesni := density(30, 30, 70, 70)
		// if math.Abs(smallDesni-allDesni) > 0.064 {
		// 	// log.Println(i)
		//
		// 	printPoints()
		// }

		// for p := range points {
		// 	if p.X == 0 {
		// 		printPoints()
		// 	}
		// }
	}

	return 0
}
