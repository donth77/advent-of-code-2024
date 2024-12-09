/*
*
--- Day 7: Bridge Repair ---
https://adventofcode.com/2024/day/7
*
*/
package day7

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var eqComboMap = make(map[int][][]string)
var eqComboMapPt2 = make(map[int][][]string)

func generateEqCombos(length int) [][]string {
	if combos, exists := eqComboMap[length]; exists {
		return combos
	}

	arr := make([][]string, 0, int(math.Pow(2, float64(length))))
	for i := 0; i < int(math.Pow(2, float64(length))); i++ {
		combo := make([]string, length)
		num := i
		for j := 0; j < length; j++ {
			if num%2 == 0 {
				combo[j] = "+"
			} else {
				combo[j] = "*"
			}
			num /= 2
		}
		arr = append(arr, combo)
	}
	eqComboMap[length] = arr
	return arr
}

func generateEqCombosPt2(length int) [][]string {
	if combos, exists := eqComboMapPt2[length]; exists {
		return combos
	}

	arr := make([][]string, 0, int(math.Pow(3, float64(length))))
	for i := 0; i < int(math.Pow(3, float64(length))); i++ {
		combo := make([]string, length)
		num := i
		for j := 0; j < length; j++ {
			switch num % 3 {
			case 0:
				combo[j] = "+"
			case 1:
				combo[j] = "*"
			case 2:
				combo[j] = "||"
			}
			num /= 3
		}
		arr = append(arr, combo)
	}
	eqComboMapPt2[length] = arr
	return arr
}

func evaluate(nums []int, ops []string, testValue int, part2Enabled bool) int {
	result := nums[0]
	for i := 0; i < len(ops)-1; i++ {
		operator := ops[i]
		nextNum := nums[i+1]
		switch operator {
		case "+":
			result += nextNum
		case "*":
			result *= nextNum
		case "||":
			if part2Enabled {
				// Convert numbers to strings, concatenate, then back to int
				combined := fmt.Sprintf("%d%d", result, nextNum)
				var err error
				result, err = strconv.Atoi(combined)
				if err != nil {
					return -1
				}
			}
		}
		if result > testValue {
			return -1
		}
	}
	return result
}

func canBeSatisfied(testValue int, nums []int, part2Enabled bool) bool {
	var eqCombos [][]string
	if part2Enabled {
		eqCombos = generateEqCombosPt2(len(nums))
	} else {
		eqCombos = generateEqCombos(len(nums))
	}

	for _, ops := range eqCombos {
		result := evaluate(nums, ops, testValue, part2Enabled)
		if result == testValue {
			return true
		}
	}
	return false
}

func ReadFile() {
	start := time.Now()

	data, err := os.ReadFile("./day07/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var part1Result, part2Result int

	for _, line := range lines {
		parts := strings.Split(line, ":")
		testValue, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Error parsing test value:", err)
			continue
		}

		numStrs := strings.Fields(strings.TrimSpace(parts[1]))
		nums := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			nums[i], err = strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				continue
			}
		}

		if canBeSatisfied(testValue, nums, false) {
			part1Result += testValue
		}
		if canBeSatisfied(testValue, nums, true) {
			part2Result += testValue
		}
	}

	fmt.Printf("Part 1:\n%d\n\nPart 2:\n%d\n", part1Result, part2Result)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %v\n", elapsed)
}
