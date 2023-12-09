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

func Part1(input []string) int {

	directions := make([]string, 0)
	for _, character := range input[0] {
		directions = append(directions, string(character))
	}

	listOfMaps := make(map[string]Directions)
	listOfStartPoints := make(map[string][]Directions)
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
			listOfStartPoints["A"] = append(listOfStartPoints["A"], point)
		}
	}

	fmt.Println("List of Start Points: ", listOfStartPoints)

	currentPoint := listOfStartPoints["A"][0]

	iterator := 0
	for !(strings.HasSuffix(currentPoint.name, "Z")) {
		for _, point := range listOfStartPoints["A"] {
			if iterator != len(directions) {
				if directions[iterator] == "L" {
					currentPoint = listOfMaps[point.left]
					count = count + 1
					iterator = iterator + 1
				} else {
					currentPoint = listOfMaps[point.right]
					count = count + 1
					iterator = iterator + 1
				}
			} else {
				iterator = 0
			}
		}
	}
	return count
}
