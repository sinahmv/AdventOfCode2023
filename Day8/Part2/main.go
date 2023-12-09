package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Directions struct {
	name  string
	left  string
	right string
}

func main() {
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day8/input.txt")
	input := string(file)
	lines := SplitLines(input)

	fmt.Println("Ergebnis Part 1:", Part1(lines))
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func Part1(input []string) []int {

	directions := make([]string, 0)
	for _, character := range input[0] {
		directions = append(directions, string(character))
	}

	listOfMaps := make(map[string]Directions)
	listOfStartPoints := make(map[string]Directions)
	count := 0

	for i := 2; i < len(input); i++ {
		line := input[i]
		line = strings.Replace(line, "(", "", -1)
		line = strings.Replace(line, ")", "", -1)
		line = strings.Replace(line, ",", "", -1)
		line = strings.Replace(line, "=", "", -1)
		splitInput := strings.Fields(line)
		foo := Directions{name: splitInput[0], left: splitInput[1], right: splitInput[2]}
		listOfMaps[splitInput[0]] = foo
	}

	for _, point := range listOfMaps {
		if strings.HasSuffix(point.name, "A") {
			listOfStartPoints[point.name] = point
		}
	}

	countOfSteps := make([]int, 0)

	iterator := 0
	for _, startpoint := range listOfStartPoints {
		for !strings.HasSuffix(startpoint.name, "Z") {
			if iterator != len(directions) {
				if directions[iterator] == "L" {
					startpoint = listOfMaps[startpoint.left]
					count = count + 1
					iterator = iterator + 1
				} else {
					startpoint = listOfMaps[startpoint.right]
					count = count + 1
					iterator = iterator + 1
				}
			} else {
				iterator = 0
			}
		}
		countOfSteps = append(countOfSteps, count)
		count = 0
	}
	return countOfSteps
}
