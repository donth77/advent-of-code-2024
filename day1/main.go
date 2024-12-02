/*
*
--- Day 1: Historian Hysteria ---
https://adventofcode.com/2024/day/1
*
*/
package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadFile() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	leftNums, rightNums := []int{}, []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		leftNum, _ := strconv.Atoi(line[0])
		rightNum, _ := strconv.Atoi(line[1])
		leftNums = append(leftNums, leftNum)
		rightNums = append(rightNums, rightNum)
	}
	sort.Ints(leftNums)
	sort.Ints(rightNums)

	totalDist := 0
	freq := make(map[int]int)
	for i, _ := range leftNums {
		totalDist += int(math.Abs(float64(leftNums[i] - rightNums[i])))
		freq[rightNums[i]] = freq[rightNums[i]] + 1

	}

	simScore := 0
	for _, num := range leftNums {
		if count, exists := freq[num]; exists {
			simScore += num * count
		}
	}

	fmt.Printf("Part 1:\n%d\n\nPart 2:\n%d\n", totalDist, simScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
