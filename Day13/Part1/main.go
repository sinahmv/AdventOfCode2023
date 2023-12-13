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

	part1 := 0
	currentGrid := make([]string, 0)

	for _, line := range lines {

		if line == "" && len(currentGrid) > 0 {
			coloums := checkLeftRight(currentGrid)
			rows := checkUpDown(currentGrid)
			result := coloums + (rows * 100)
			part1 = part1 + result
			currentGrid = make([]string, 0)
		} else {
			currentGrid = append(currentGrid, line)
		}
	}

	return part1
}

func checkLeftRight(grid []string) int {

	for counter := 0; counter < len(grid[0])-1; counter++ {
		diff := false

		for line := 0; line < len(grid); line++ {
			for offset := 0; ; offset++ {
				
				left := counter - offset
				right := counter + offset + 1
				if left < 0 || right >= len(grid[0]) {
					break
				}
				if grid[line][left] != grid[line][right] {
					diff = true
				}
			}
		}
		if !diff {
			return counter + 1
		}
	}
	return 0
}

func checkUpDown(grid []string) int {

	for counter := 0; counter < len(grid)-1; counter++ {
		diff := false

		for col := 0; col < len(grid[0]); col++ {
			for offset := 0; ; offset++ {

				lowRow := counter - offset
				highRow := counter + offset + 1
				if lowRow < 0 || highRow >= len(grid) {
					break
				}
				if grid[lowRow][col] != grid[highRow][col] {
					diff = true
				}
			}
		}
		if !diff {
			return counter + 1
		}
	}
	return 0
}

func SplitLine(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
