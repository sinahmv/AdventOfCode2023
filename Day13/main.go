package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day13/input.txt")
	input := string(file)
	lines := SplitLine(input)
	fmt.Println("Ergebnis Part 1: ", Part1(lines))
}

func Part1(lines []string) int {

	output := make ([]int,0)
	part1 := 0

	currentGrid := make([]string,0)

	for _, line := range lines {

		if line == "" && len(currentGrid) > 0 {
			left := compareLeftRight(currentGrid)
			up := compareUpDown(currentGrid)
			result := left + up*100
			output = append(output, result)
			currentGrid = make([]string, 0)
		} else {
			currentGrid = append(currentGrid, line)
		}
	}

	for _,result := range output {
		part1 = part1 + result
	}
	return part1
}

func compareLeftRight(grid []string) int {

	output := 0

	for i := 1; i < len(grid[0]); i++ {

		count := 0
		left := i - 1
		right := i
		sRange := min(i, len(grid[0])-i)

		for j := 0; j < sRange; j++ {
			for _, line := range grid {
				charLeft := left - j
				charRight := right + j
				if line[charLeft] == line[charRight] {
					count++
				}
			}
		}
		output = i
	}
	return output
}

func compareUpDown(lines []string) int {

	output := 0

	for i := 0; i < len(lines)-1; i++ {
		line := i
		lineDown := i + 1
		count := 0

		for j := 0; j < len(lines[line]); j++ {
			string := lines[line]
			stringDown := lines[lineDown]
			if string == stringDown {
				count++
			}
		}
		output = i
	}
	return output
}

func SplitLine(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
