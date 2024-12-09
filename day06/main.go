/*
*
--- Day 6: Guard Gallivant ---
https://adventofcode.com/2024/day/6
*
*/
package day6

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Direction struct {
	dr, dc int
}

var dirs = map[int]Direction{
	0: {-1, 0}, // up
	1: {0, 1},  // right
	2: {1, 0},  // down
	3: {0, -1}, // left
}

type Point struct {
	row, col int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.row, p.col)
}

type State struct {
	point Point
	dir   int
}

func (s State) String() string {
	return fmt.Sprintf("%d,%d,%d", s.point.row, s.point.col, s.dir)
}

func inBounds(row, col, rows, cols int) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols
}

type SimulationResult struct {
	isLoop  bool
	visited map[string]struct{}
}

func simulateGuard(startRow, startCol int, testGrid [][]string, rows, cols int) SimulationResult {
	movingDir := 0
	gr, gc := startRow, startCol
	stateSet := make(map[string]struct{})
	visited := make(map[string]struct{})
	maxSteps := rows * cols * 4
	steps := 0

	visited[Point{gr, gc}.String()] = struct{}{}

	for inBounds(gr, gc, rows, cols) && steps < maxSteps {
		state := State{Point{gr, gc}, movingDir}
		stateKey := state.String()

		if _, exists := stateSet[stateKey]; exists {
			return SimulationResult{true, visited}
		}
		stateSet[stateKey] = struct{}{}

		dir := dirs[movingDir]
		nextRow := gr + dir.dr
		nextCol := gc + dir.dc

		if !inBounds(nextRow, nextCol, rows, cols) {
			break
		}

		if testGrid[nextRow][nextCol] == "#" || testGrid[nextRow][nextCol] == "O" {
			movingDir = (movingDir + 1) % 4
		} else {
			gr = nextRow
			gc = nextCol
			visited[Point{gr, gc}.String()] = struct{}{}
		}
		steps++
	}

	return SimulationResult{false, visited}
}

func copyGrid(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid))
	for i := range grid {
		newGrid[i] = make([]string, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}
func ReadFile() {
	start := time.Now()

	data, err := os.ReadFile("./day06/input.txt")
	if err != nil {
		panic(err)
	}

	// Parse grid
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	rows, cols := len(grid), len(grid[0])
	guardRow, guardCol := 0, 0

	// Find guard position
	for r := 0; r < len(grid); r++ {
		if strings.Contains(strings.Join(grid[r], ""), "^") {
			guardRow = r
			for c, ch := range grid[r] {
				if ch == "^" {
					guardCol = c
					break
				}
			}
			break
		}
	}

	// Part 1
	part1Result := simulateGuard(guardRow, guardCol, grid, rows, cols)
	fmt.Printf("Part 1:\n%d\n", len(part1Result.visited))

	// Part 2
	loops := make(map[string]struct{})
	testGrid := copyGrid(grid)

	for step := range part1Result.visited {
		var r, c int
		fmt.Sscanf(step, "%d,%d", &r, &c)

		if grid[r][c] != "." {
			continue
		}

		testGrid[r][c] = "O"
		result := simulateGuard(guardRow, guardCol, testGrid, rows, cols)
		if result.isLoop {
			loops[Point{r, c}.String()] = struct{}{}
		}
		testGrid[r][c] = "."
	}

	fmt.Printf("\nPart 2:\n%d\n", len(loops))

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %v\n", elapsed)
}
