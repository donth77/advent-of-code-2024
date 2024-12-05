/*
*
--- Day 2: Red-Nosed Reports ---
https://adventofcode.com/2024/day/2
*
*/
package day2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadFile() {
	file, err := os.Open("./day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	isSequenceSafe := func(nums []int) bool {
		if len(nums) < 2 {
			return true
		}

		var isIncreasing *bool
		for i := 1; i < len(nums); i++ {
			diff := nums[i] - nums[i-1]
			absDiff := int(math.Abs(float64(diff)))

			if absDiff < 1 || absDiff > 3 {
				return false
			}

			diffIsPos := diff > 0

			if isIncreasing == nil {
				isIncreasing = &diffIsPos
				continue
			}

			if diffIsPos != *isIncreasing {
				return false
			}
		}

		return true
	}

	isSequenceSafePt2 := func(nums []int) bool {
		if isSequenceSafe(nums) {
			return true
		}

		// Remove each and check if safe
		for i := 0; i < len(nums); i++ {
			// Create new slice without element at index i
			modified := make([]int, 0, len(nums)-1)
			modified = append(modified, nums[:i]...)
			modified = append(modified, nums[i+1:]...)

			if isSequenceSafe(modified) {
				return true
			}
		}

		return false
	}

	scanner := bufio.NewScanner(file)
	numSafe := 0
	numSafePt2 := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		nums := make([]int, len(line))

		for i, str := range line {
			num, _ := strconv.Atoi(str)
			nums[i] = num
		}
		isSafe := isSequenceSafe(nums)
		isSafePt2 := isSequenceSafePt2(nums)

		if isSafe {
			numSafe++
		}
		if isSafePt2 {
			numSafePt2++
		}
	}

	fmt.Printf("Part 1:\n%d\n\nPart 2:\n%d\n", numSafe, numSafePt2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
