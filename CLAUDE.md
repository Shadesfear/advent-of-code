# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Structure

Advent of Code solutions organized by year and language:
- `2024/go/dayXX/` and `2025/go/dayXX/` - Go solutions (current primary language)
- `2022/golang/` and `2022/rust-aoc/` - Mixed Go and Rust from 2022
- `2021/` - Common Lisp solutions
- `lib/go/` - Shared Go library (`github.com/shadesfear/aoc-lib-go`)
- `YYYY/inputs/dayXX.txt` - Puzzle inputs per year

## Commands

```bash
# Run tests for a specific day
go test -v ./2025/go/day03/

# Run a single test
go test -v -run TestPart1 ./2025/go/day03/

# Run solution
go run ./2025/go/day03/

# Initialize new day (from go-day-init directory)
cd go-day-init && go run . && cd -
# Then move generated files to appropriate day directory
```

## Shared Library (lib/go)

The library is imported as `github.com/shadesfear/aoc-lib-go`. Key packages:

- `files`: Input parsing (`ReadInputLines`, `ParseInputToGrid`, `ReadDayInput`)
- `datastructures`: `Point`, `Dir4`, `Set`, `Stack`, `Queue`, `PriorityQueue`
- `str`: String utilities (`ToInt`, `ToInt64`, `SplitLines`, `PrettyPrintGrid`)
- `math`: `GCD`, `LCM`, `Abs`, `Pow`, `Permutations`

## Solution Pattern

Each day follows this structure:
- `main.go` with `solvePart1(lines []string) int` and `solvePart2(lines []string) int`
- `main_test.go` with `TestPart1` and `TestPart2` using example input
- Input read from `../../inputs/dayXX.txt` relative to day directory

## Point/Grid Conventions

- Point: `X` is column, `Y` is row
- Dir4 directions: `Up` (0,-1), `Right` (1,0), `Down` (0,1), `Left` (-1,0)
- Grid bounds check: `point.InBounds(rows, cols)` where rows = height, cols = width
