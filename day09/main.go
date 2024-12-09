/*
*
--- Day 9: Disk Fragmenter ---
https://adventofcode.com/2024/day/9
*
*/
package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile() {
	data, err := os.ReadFile("./day09/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var disk []int
	id := 0

	content := strings.TrimSpace(string(data))
	for i, char := range content {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Printf("Error parsing number at position %d: %v\n", i, err)
			return
		}

		for j := 0; j < num; j++ {
			if i%2 == 0 {
				disk = append(disk, id)
			} else {
				disk = append(disk, -1)
			}
		}
		if i%2 == 0 {
			id++
		}
	}

	// Create copy for part 2
	disk2 := make([]int, len(disk))
	copy(disk2, disk)

	// -- Part 1 --
	i, j := 0, len(disk)-1
	for i < j {
		for i < len(disk) && disk[i] != -1 {
			i++
		}
		for j >= 0 && disk[j] == -1 {
			j--
		}
		if i < j {
			disk[i], disk[j] = disk[j], disk[i]
		}
	}

	checksum := 0
	for i, val := range disk {
		if val > 0 {
			checksum += i * val
		}
	}

	// -- Part 2 --
	for fileId := id - 1; fileId >= 0; fileId-- {
		// Find file location and size
		fStart := -1
		fSize := 0
		for i, val := range disk2 {
			if val == fileId {
				if fStart == -1 {
					fStart = i
				}
				fSize++
			}
		}

		if fStart == -1 {
			continue
		}

		// Find gap
		bestGapStart := -1
		pos := 0
		for pos < fStart {
			// Skip non-empty spaces
			for pos < fStart && disk2[pos] != -1 {
				pos++
			}
			if pos >= fStart {
				break
			}

			// Measure gap size
			gapSize := 0
			gapStart := pos
			for pos < fStart && disk2[pos] == -1 {
				gapSize++
				pos++
			}

			// Update gap 
			if fSize <= gapSize && (bestGapStart == -1 || gapStart < bestGapStart) {
				bestGapStart = gapStart
			}
		}

		// Move file if valid gap found
		if bestGapStart != -1 {
			for i := 0; i < fSize; i++ {
				disk2[bestGapStart+i] = fileId // Place file in gap
				disk2[fStart+i] = -1           // Clear old location
			}
		}
	}

	// Calculate checksum for part 2
	checksumPt2 := 0
	for i, val := range disk2 {
		if val > 0 {
			checksumPt2 += i * val
		}
	}

	fmt.Printf("Part 1:\n%d\n\nPart 2:\n%d\n", checksum, checksumPt2)
}
