/*
*
--- Day 18:  ---
https://adventofcode.com/2024/day/18
*
*/
package day18

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile() {
	file, err := os.Open("./day18/sample-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		fmt.Println(line)
	}

	fmt.Printf("Part 1:\n%d\n\nPart 2:\n%d\n", 0, 0)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
