/*
*
--- Day 4: Ceres Search ---
https://adventofcode.com/2024/day/4
*
*/
package day4

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Direction struct to represent movement vectors
type Direction struct {
	row, col int
}

// Count "XMAS" in all directions
func CountXMAS(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	// Define all possible directions
	dirs := []Direction{
		{0, 1},   // right
		{1, 0},   // down
		{1, 1},   // diagonal down-right
		{-1, 1},  // diagonal up-right
		{0, -1},  // left
		{-1, 0},  // up
		{-1, -1}, // diagonal up-left
		{1, -1},  // diagonal down-left
	}

	// isValid checks if a position is within grid bounds
	isValid := func(row, col int) bool {
		return row >= 0 && row < rows && col >= 0 && col < cols
	}

	// checkWord checks for "XMAS" starting from a position in a direction
	checkWord := func(row, col int, dir Direction) bool {
		word := "XMAS"
		for i := 0; i < len(word); i++ {
			newRow := row + (dir.row * i)
			newCol := col + (dir.col * i)
			if !isValid(newRow, newCol) || grid[newRow][newCol] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	// Check each position in the grid
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 'X' {
				for _, dir := range dirs {
					if checkWord(row, col, dir) {
						count++
					}
				}
			}
		}
	}

	return count
}

// Count X-shaped patterns of "MAS"/"SAM"
func CountXMASPatterns(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	// isValid checks if a position is within grid bounds
	isValid := func(row, col int) bool {
		return row >= 0 && row < rows && col >= 0 && col < cols
	}

	// getDiagonalString gets a 3-character string in a diagonal direction
	getDiagonalString := func(row, col int, dirRow, dirCol int) string {
		if !isValid(row, col) ||
			!isValid(row+dirRow, col+dirCol) ||
			!isValid(row+2*dirRow, col+2*dirCol) {
			return ""
		}
		return string([]rune{
			grid[row][col],
			grid[row+dirRow][col+dirCol],
			grid[row+2*dirRow][col+2*dirCol],
		})
	}

	// isMAS checks if a string is "MAS" or "SAM"
	isMAS := func(str string) bool {
		return str == "MAS" || str == "SAM"
	}

	// checkXPattern checks for X-shaped pattern at a position
	checkXPattern := func(row, col int) bool {
		strings := []string{
			getDiagonalString(row, col, 1, 1),    // Top-left to bottom-right
			getDiagonalString(row, col+2, 1, -1), // Top-right to bottom-left
		}
		return isMAS(strings[0]) && isMAS(strings[1])
	}

	// Check each valid starting position
	for row := 0; row < rows-2; row++ {
		for col := 0; col < cols-2; col++ {
			if checkXPattern(row, col) {
				count++
			}
		}
	}

	return count
}

func ReadFile() {
	file, err := os.ReadFile("./day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileStr := string(file)

	grid := make([][]rune, 0)
	for _, line := range strings.Split(strings.TrimSpace(fileStr), "\n") {
		grid = append(grid, []rune(line))
	}

	fmt.Printf("Part 1:\n%d\n\nPart 2:\n%d\n",
		CountXMAS(grid),
		CountXMASPatterns(grid))
}
