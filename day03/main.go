/*
*
--- Day 3: Mull It Over  ---
https://adventofcode.com/2024/day/3
*
*/
package day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func processMulInstructions(input string, part1 bool) int {
	part1Result := 0
	part2Result := 0
	enabled := true

	// Go regex syntax is slightly different from JavaScript
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	// Find all matches in the input string
	matches := regex.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		instr := match[0]

		if instr == "do()" {
			enabled = true
		} else if instr == "don't()" {
			enabled = false
		} else {
			// Convert string numbers to integers
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			part1Result += num1 * num2
			if enabled {
				part2Result += num1 * num2
			}
		}
	}

	if part1 {
		return part1Result
	}
	return part2Result
}

func ReadFile() {
	file, err := os.ReadFile("./day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileStr := string(file)

	fmt.Printf("Part 1:\n%d\n\nPart 2:\n%d\n",
		processMulInstructions(fileStr, true),
		processMulInstructions(fileStr, false))
}
