package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day9/input.txt")
	input := string(file)
	lines := SplitLines(input)

	fmt.Println("Ergebnis Part1: ", Part1(lines))
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func Part1(input []string) int {

	count := 0

	for _, line := range input {

		splitInput := strings.Fields(line)
		newLine := make([]int, 0)
		for _, field := range splitInput {
			numb, _ := strconv.Atoi(string(field))
			newLine = append(newLine, numb)
		}
		fmt.Println("Line: ", newLine)

		numb := generateNewLine(newLine)
		count = count + numb
	}
	return count
}

func generateNewLine(line []int) int {

	if countOccurrences(line, 0, len(line)) {
		fmt.Println("NewValue: ", 0)
		return 0
	}

	nextLine := make([]int, 0)
	diff := 0

	for i := 0; i < len(line)-1; i++ {
		firstNumb := line[i]
		secondNumb := line[i+1]
		diff = secondNumb - (firstNumb)
		nextLine = append(nextLine, diff)
	}
	
	fmt.Println("NextLine: FOO", nextLine)

	newValueForLine := generateNewLine(nextLine) + line[len(line)-1]
	fmt.Println("NewValue: ", newValueForLine)
	return newValueForLine
}

func countOccurrences(m []int, value int, len int) bool {
	foo := false
	count := 0
	for _, v := range m {
		if v == value {
			count++
			if count == len {
				foo = true
			}
		}
	}
	return foo
}
