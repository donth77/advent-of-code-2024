/*
*
--- Day 8: Resonant Collinearity ---
https://adventofcode.com/2024/day/8
*
*/
package day8

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func distance(p1, p2 Point) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return dx*dx + dy*dy
}

func areCollinear(p1, p2, p3 Point) bool {
	// Collinear if the area of triangle formed is 0
	// Area = x1(y2 - y3) + x2(y3 - y1) + x3(y1 - y2)
	area := p1.x*(p2.y-p3.y) + p2.x*(p3.y-p1.y) + p3.x*(p1.y-p2.y)
	return area == 0
}

func isAntinode(ant1, ant2, point Point, part2 bool) bool {
	if !areCollinear(ant1, ant2, point) {
		return false
	}
	if part2 {
		return true
	}

	dist1 := distance(ant1, point)
	dist2 := distance(ant2, point)
	return dist1 == 4*dist2 || dist2 == 4*dist1
}

func findAntinodes(grid [][]string, part2 bool) map[string]struct{} {
	antinodes := make(map[string]struct{})

	// Create frequency map of antenna positions
	freq := make(map[string][]Point)
	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] != "." {
				antenna := grid[y][x]
				freq[antenna] = append(freq[antenna], Point{x, y})
			}
		}
	}

	// Check each pair of antennas
	for _, antennas := range freq {
		for _, antenna1 := range antennas {
			for _, antenna2 := range antennas {
				if antenna1 == antenna2 {
					continue
				}
				// Check each point in the grid
				for y := range grid {
					for x := range grid[0] {
						point := Point{x, y}
						if isAntinode(antenna1, antenna2, point, part2) {
							antinodes[point.String()] = struct{}{}
						}
					}
				}
			}
		}
	}
	return antinodes
}

func ReadFile() {
	data, err := os.ReadFile("./day08/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse grid
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	antinodes := findAntinodes(grid, false)
	antinodesPt2 := findAntinodes(grid, true)

	fmt.Printf("Part 1:\n%d\n\nPart 2:\n%d\n", len(antinodes), len(antinodesPt2))
}
