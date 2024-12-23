package main

import (
	// "fmt"
	"container/heap"
	"log"

	"github.com/shadesfear/aoc-lib-go/files"

	// "github.com/shadesfear/aoc-lib-go/math"

	ds "github.com/shadesfear/aoc-lib-go/datastructures"
)

type grid [][]rune

var GRID grid

func (gr grid) charAt(p Point) rune {
	return gr[p.Y][p.X]
}

func (gr grid) width() int {
	return len(gr[0])
}

func (gr grid) height() int {
	return len(gr)
}

type Point = ds.Point

func main() {
	lines, err := files.ReadInputFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input")
	}

	grid, err := files.ParseInputToGrid(lines)
	if err != nil {
		log.Fatal(err)
	}

	GRID = grid
	part1 := solvePart1()
	log.Println(part1)

	part2 := solvePart2()
	log.Println(part2)
}

func find(c rune) Point {
	start := ds.NewPoint(0, 0)
	for y, row := range GRID {
		for x, char := range row {
			if char == c {
				start = ds.NewPoint(x, y)
			}
		}
	}

	return start
}

type Pointy struct {
	Pos Point
	Dir ds.Dir4
}

func Dijkstra(start, end Point) (int, map[Pointy]int) {
	distances := map[Pointy]int{}
	poss := ds.NewSet[Pointy]()
	startDir := ds.Right
	distances[Pointy{start, startDir}] = 0

	for {

		var minPointy Pointy
		minDist := 100000000

		ran := false
		for p, dist := range distances {
			if poss.Contains(p) {
				continue
			}
			if dist < minDist {
				minPointy = p
				minDist = dist
			}

			ran = true
		}

		if !ran {
			break
		}

		dir := minPointy.Dir
		for _, v := range []ds.Dir4{dir, dir.RotateCCW(), dir.RotateCW()} {
			next := minPointy.Pos.MoveDir(v)
			if !next.InBounds(GRID.width(), GRID.height()) {
				continue
			}
			if GRID.charAt(next) == '#' {
				continue
			}
			nextDist := minDist + 1
			if v != dir {
				nextDist += 1000
			}
			nextPointy := Pointy{next, v}
			distances[nextPointy] = nextDist

		}

		poss.Add(minPointy)

	}
	res := 10000000

	for _, r := range []int{
		distances[Pointy{end, ds.Right}],
		distances[Pointy{end, ds.Up}],
		distances[Pointy{end, ds.Left}],
		distances[Pointy{end, ds.Down}],
	} {
		if r < res && r > 0 {
			res = r
		}
	}
	return res, distances
}

type Adj struct {
	p    Pointy
	cost int
}

func adjs(p Pointy) []Adj {
	adj := []Adj{
		{Pointy{p.Pos, p.Dir.RotateCW()}, 1000},
		{Pointy{p.Pos, p.Dir.RotateCCW()}, 1000},
	}
	next := p.Pos.MoveDir(p.Dir)
	if next.InBounds(GRID.height(), GRID.width()) && GRID.charAt(next) != '#' {
		adj = append(adj, Adj{Pointy{next, p.Dir}, 1})
	}
	return adj
}

func DijkstraPQ(start, end Point) int {
	pq := ds.PriorityQueue{}
	heap.Push(
		&pq,
		&ds.Item{
			Value:    Pointy{start, ds.Right},
			Priority: 0,
		},
	)

	heap.Init(&pq)

	for pq.Len() > 0 {
		i := heap.Pop(&pq)
		dist := i.(ds.Item).Priority
		cur := i.(ds.Item).Value.(Pointy)

		if cur.Pos.Equal(end) {
			break
		}

	}

	return 0
}

func findOptimalPathPoints(start, end Point) map[Point]bool {
	// First run Dijkstra's to get minimum cost and distances
	minCost, distances := Dijkstra(start, end)

	// Track all points that are part of an optimal path
	optimalPoints := make(map[Point]bool)

	// For each point and direction in the grid, check if it could be part of an optimal path
	for y := 0; y < GRID.height(); y++ {
		for x := 0; x < GRID.width(); x++ {
			pos := Point{X: x, Y: y}
			if GRID.charAt(pos) == '#' {
				continue
			}

			// For each direction at this point
			for _, dir := range []ds.Dir4{ds.Right, ds.Up, ds.Left, ds.Down} {
				pointy := Pointy{pos, dir}
				costToHere, existsHere := distances[pointy]
				if !existsHere {
					continue
				}

				// Try moving in each possible direction
				for _, nextDir := range []ds.Dir4{dir, dir.RotateCW(), dir.RotateCCW()} {
					nextPos := pos.MoveDir(nextDir)
					if !nextPos.InBounds(GRID.width(), GRID.height()) {
						continue
					}
					if GRID.charAt(nextPos) == '#' {
						continue
					}

					// Calculate cost to move to next position
					turnCost := 0
					if nextDir != dir {
						turnCost = 1000
					}

					// Check if moving through this point could be part of an optimal path
					for _, finalDir := range []ds.Dir4{ds.Right, ds.Up, ds.Left, ds.Down} {
						if endCost, exists := distances[Pointy{end, finalDir}]; exists {
							if endCost == minCost {
								nextPointy := Pointy{nextPos, nextDir}
								if nextCost, existsNext := distances[nextPointy]; existsNext {
									if costToHere+1+turnCost == nextCost {
										optimalPoints[pos] = true
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// Also mark the end point
	optimalPoints[end] = true

	return optimalPoints
}

func solvePart1() int {
	res := 0

	start := find('S')
	end := find('E')
	res, _ = Dijkstra(start, end)

	return res
}

func solvePart2() int {
	start := find('S')
	end := find('E')

	res, _ := Dijkstra(start, end)

	o := findOptimalPathPoints(start, end)
	log.Println("")
	log.Println(o)
	log.Println(len(o))

	// log.Println(from)

	// seen := ds.NewSet[Point]()
	//
	// stack := []Point{end} // Start from end point
	//
	// log.Printf("Starting at end: %v\n", end)
	// log.Printf("from[end] contains: %v\n", from[end])
	//
	// processed := ds.NewSet[string]()
	//
	// for len(stack) > 0 {
	// 	current := stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]
	// 	seen.Add(current)
	//
	// 	// Add all points that can reach current
	// 	for _, prev := range from[current] {
	// 		edge := fmt.Sprintf("%v->%v", current, prev)
	// 		if !processed.Contains(edge) {
	// 			processed.Add(edge)
	// 			stack = append(stack, prev)
	// 		}
	// 	}
	// }
	// log.Printf("Final seen points: %v\n", seen.Items)
	// for !st.IsEmpty() {
	// }

	return res // seen.Len()
}
