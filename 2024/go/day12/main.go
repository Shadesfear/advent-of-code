package main

import (
	"log"
	"slices"

	"github.com/shadesfear/aoc-lib-go/datastructures"
	ds "github.com/shadesfear/aoc-lib-go/datastructures"
	"github.com/shadesfear/aoc-lib-go/files"
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

func isBoundary(toLookFor byte, me ds.Point, lines []string) bool {
	counter := 0

	var seen []ds.Point
	//nabos := me.Neighbours()
	for _, nabo := range me.Neighbours() {
		nabo.
		if slices.Contains(seen, nabo) {
			continue
		}
		if !nabo.InBounds(len(lines), len(lines[0])) {
			counter++
			continue
		}
		if lines[nabo.Y][nabo.X] == toLookFor {
			continue
		}
		counter++
	}

	if counter > 0 {
		return true
	}

	return false
}

var dfsseen map[ds.Point]int = map[ds.Point]int{}

func dfs(lines []string, point ds.Point, t byte, typeCounter int) {
	_, contains := dfsseen[point]
	if contains {
		return
	}
	if lines[point.Y][point.X] == t {
		dfsseen[point] = typeCounter
		for _, nabo := range point.Neighbours() {
			if !nabo.InBounds(len(lines), len(lines[0])) {
				continue
			}
			dfs(lines, nabo, t, typeCounter)
		}
	}
}

func bfs(lines []string, start ds.Point) []ds.Point {
	seen := ds.NewSet[ds.Point]()
	want := ds.NewSet[ds.Point]()

	block := lines[start.Y][start.X]

	q := ds.NewQueue[ds.Point]()
	q.Enqueue(start)

	for !q.IsEmpty() {
		cur, ok := q.Dequeue()
		if !ok {
			log.Fatal("Not okay")
		}

		nabos := cur.Neighbours()

		for _, nabo := range nabos {
			// log.Printf("Looking at nabo: %s", nabo)
			if !nabo.InBounds(len(lines), len(lines[0])) {
				continue
			}
			if seen.Contains(nabo) {
				continue
			}
			naboByte := lines[nabo.Y][nabo.X]
			if naboByte != block {
				continue
			}
			if !isBoundary(block, nabo, lines) {
				continue
			}

			q.Enqueue(nabo)
		}
		seen.Add(cur)
		want.Add(cur)

	}

	return want.ToSlice()
}

func findPeri(lines []string, block []ds.Point) int {
	weAre := lines[block[0].Y][block[0].X]
	seen := []ds.Point{}
	peri := 0

	for _, field := range block {
		if slices.Contains(seen, field) {
			continue
		}

		nabos := field.Neighbours()
		for _, nabo := range nabos {
			if slices.Contains(seen, nabo) {
				continue
			}
			if !nabo.InBounds(len(lines), len(lines[0])) {
				peri += 1
				continue
			}
			if lines[nabo.Y][nabo.X] == weAre {
				continue
			}
			peri += 1
		}
	}
	return peri
}

func findArea(lines []string, block []ds.Point) int {
	area := 0
	weAre := lines[block[0].Y][block[0].X]

	// log.Printf("Finding area of block: %s", block)

	for y, row := range lines {
		for x := range row {
			point := ds.NewPoint(x, y)

			if slices.Contains(block, point) {
				continue
			}
			if isInteriorPoint(lines, block, point) && lines[point.Y][point.X] == weAre {
				// log.Printf("We are interior point and buddies: %s", point)
				// log.Println(point)
				// log.Println(area)
				area++
			}
		}
	}
	// log.Printf("Area: %d", area)
	return area + len(block)
}

func interiorPoints(block []ds.Point) int {
	A := int(datastructures.Shoelace(block))
	ip := int(A) - (len(block) / 2) + 1
	if ip < 0 {
		return 0
	}
	return ip
}

func isInteriorPoint(lines []string, block []ds.Point, cand ds.Point) bool {
	crosses := 0
	cur := cand

	if len(block) == 1 {
		return false
	}

	dir := ds.NewPoint(1, 0)
	for {
		next := cur.Move(dir.X, dir.Y)
		if !next.InBounds(len(lines), len(lines[0])) {
			break
		}
		if slices.Contains(block, next) && !slices.Contains(block, cur) {
			crosses++
		}
		cur = next
	}
	if crosses == 0 {
		return false
	}

	if crosses%2 == 0 {
		return false
	}
	return true
}

func allInterior(lines []string, block []ds.Point) []ds.Point {
	interior := []ds.Point{}
	candByte := lines[block[0].Y][block[0].X]

	if len(block) == 1 {
		return interior
	}

	for y := range lines {
		for x := range lines[0] {
			cand := ds.NewPoint(x, y)
			if slices.Contains(block, cand) {
				continue
			}
			if candByte != lines[y][x] {
				continue
			}
			if isInteriorPoint(lines, block, cand) {
				interior = append(interior, cand)
			}

		}
	}
	return interior
}

func solvePart1(lines []string) int {
	res := 0

	counter := 0
	for y, row := range lines {
		for x := range row {
			p := ds.NewPoint(x, y)
			_, ok := dfsseen[p]
			if !ok {
				dfs(lines, p, lines[p.Y][p.X], counter)
				counter++
			}
		}
	}
	// log.Println(dfsseen)

	// visited := ds.NewSet[ds.Point]()
	//
	// blocks := []ds.Set[ds.Point]{}
	//
	// for y, row := range lines {
	// 	for x := range row {
	// 		here := ds.NewPoint(x, y)
	// 		if visited.Contains(here) {
	// 			continue
	// 		}
	//
	// 		farm := bfs(lines, here)
	// 		interior := allInterior(lines, farm)
	//
	// 		blockSet := ds.NewSet[ds.Point]()
	//
	// 		blockSet.AddM(farm)
	// 		blockSet.AddM(interior)
	//
	// 		blocks = append(blocks, *blockSet)
	//
	// 		visited.AddM(farm)
	// 		visited.AddM(interior)
	//
	// 	}
	// }
	//
	// // log.Println(blocks, len(blocks))
	//
	blocks := map[int]ds.Set[ds.Point]{}
	for k, v := range dfsseen {
		_, ok := blocks[v]

		if !ok {
			blocks[v] = *ds.NewSet[ds.Point]()
		}
		s := blocks[v]
		s.Add(k)

	}

	// log.Println("blocks")
	// log.Println(blocks)

	for _, block := range blocks {
		// s := block.ToSlice()
		// log.Println(string(lines[s[0].Y][s[0].X]))

		// log.Println(block.ToSlice())
		area := len(block.ToSlice())

		peri := findPeri(lines, block.ToSlice())
		// log.Printf("area: %d, peri: %d", area, peri)
		res += peri * area
	}

	// log.Println(res)

	return res
}

func solvePart2(lines []string) int {
	res := 0

	return res
}
