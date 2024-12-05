/*
*
--- Day 5:  ---
https://adventofcode.com/2024/day/5
*
*/
package day5

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type InputData struct {
	rules   map[int]map[int]bool // Using map[int]bool as a Set equivalent
	updates [][]int
}

func ParseInput(input string) InputData {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	rulesSection, updatesSection := sections[0], sections[1]

	// Initialize rules map
	rules := make(map[int]map[int]bool)

	// Parse rules
	for _, rule := range strings.Split(rulesSection, "\n") {
		parts := strings.Split(rule, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])

		if _, exists := rules[after]; !exists {
			rules[after] = make(map[int]bool)
		}
		rules[after][before] = true
	}

	// Parse updates
	var updates [][]int
	for _, line := range strings.Split(updatesSection, "\n") {
		var update []int
		for _, num := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(num)
			update = append(update, n)
		}
		updates = append(updates, update)
	}

	return InputData{rules: rules, updates: updates}
}

func IsValidOrder(update []int, rules map[int]map[int]bool) bool {
	for i, currentPage := range update {
		if requiredPages, exists := rules[currentPage]; exists {
			previousPages := make(map[int]bool)
			for _, page := range update[:i] {
				previousPages[page] = true
			}

			for requiredPage := range requiredPages {
				// Check if required page exists in update but not in previous pages
				exists := false
				for _, page := range update {
					if page == requiredPage {
						exists = true
						break
					}
				}
				if exists && !previousPages[requiredPage] {
					return false
				}
			}
		}
	}
	return true
}

// Creates a adjacency list
func CreateGraph(pages []int, rules map[int]map[int]bool) map[int]map[int]bool {
	graph := make(map[int]map[int]bool)

	for _, page := range pages {
		graph[page] = make(map[int]bool)
	}

	// Add edges based on rules
	for _, page := range pages {
		if dependencies, exists := rules[page]; exists {
			for dep := range dependencies {
				// Check if dep exists in pages
				exists := false
				for _, p := range pages {
					if p == dep {
						exists = true
						break
					}
				}
				if exists {
					graph[dep][page] = true
				}
			}
		}
	}
	return graph
}

// Topological sort on the pages based on rules
func TopologicalSort(pages []int, rules map[int]map[int]bool) []int {
	graph := CreateGraph(pages, rules)
	inDegree := make(map[int]int)

	for _, page := range pages {
		inDegree[page] = 0
	}

	// Calculate in-degrees
	for _, dependencies := range graph {
		for dep := range dependencies {
			inDegree[dep]++
		}
	}

	// Initialize queue with nodes having no dependencies
	var queue []int
	for _, page := range pages {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	var result []int
	for len(queue) > 0 {
		// Pop from queue
		page := queue[0]
		queue = queue[1:]
		result = append(result, page)

		// Update neighbors
		for neighbor := range graph[page] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

func GetMiddleNum(arr []int) int {
	return arr[len(arr)/2]
}

func ReadFile() {
	file, err := os.ReadFile("./day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileStr := string(file)

	inputData := ParseInput(fileStr)
	rules, updates := inputData.rules, inputData.updates

	// --- Part 1 ---
	part1Result := 0
	for _, update := range updates {
		if IsValidOrder(update, rules) {
			part1Result += GetMiddleNum(update)
		}
	}

	// --- Part 2 ---
	part2Result := 0
	for _, update := range updates {
		if !IsValidOrder(update, rules) {
			sorted := TopologicalSort(update, rules)
			part2Result += GetMiddleNum(sorted)
		}
	}

	fmt.Printf("Part 1:\n%d\n\nPart 2:\n%d\n", part1Result, part2Result)
}
