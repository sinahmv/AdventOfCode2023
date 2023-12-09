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

		numb := generateNewLine(newLine)
		count = count + numb
	}
	return count
}

func generateNewLine(line []int) int {

	if countOccurrences(line, 0, len(line)) {
		return 0
	}

	nextLine := make([]int, 0)
	diff := 0

	for i := len(line) - 1; i > 0; i-- {
		firstNumb := line[i]
		secondNumb := line[i-1]
		diff = firstNumb - secondNumb
		nextLine = insert(nextLine, 0, diff)
	}

	newValueForLine := line[0] - generateNewLine(nextLine) 
	return newValueForLine
}

func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
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
