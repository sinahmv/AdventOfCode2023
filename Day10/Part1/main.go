package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day10/input.txt")
	input := string(file)
	lines := SplitLine(input)
	matrix := SplitLines(lines)
	fmt.Println("Ergebnis Part 1: ", Part1(matrix))
}

func Part1(matrix [][]string) int {

	count := 1
	startingPoint := make([]int, 0)

	directions := map[string]map[string][]string{
		"|": {"oben": []string{"1", "0", "oben"}, "unten": []string{"-1", "0", "unten"}},
		"-": {"links": []string{"0", "1", "links"}, "rechts": []string{"0", "-1", "rechts"}},
		"L": {"oben": []string{"0", "1", "links"}, "rechts": []string{"-1", "0", "unten"}},
		"J": {"oben": []string{"0", "-1", "rechts"}, "links": []string{"-1", "0", "unten"}},
		"7": {"unten": []string{"0", "-1", "rechts"}, "links": []string{"1", "0", "oben"}},
		"F": {"unten": []string{"0", "1", "links"}, "rechts": []string{"1", "0", "oben"}},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "S" {
				startingPoint = append(startingPoint, i)
				startingPoint = append(startingPoint, j)
			}
		}
	}

	nextCoordinates, direction := FindConnectingPipe(matrix, startingPoint)
	nextPoint := matrix[nextCoordinates[0]][nextCoordinates[1]]

	for true {

		nextPoint = matrix[nextCoordinates[0]][nextCoordinates[1]]
		if nextPoint == "S" {
			break
		}
		newDirection := directions[nextPoint][direction]
		newX, _ := strconv.Atoi(newDirection[1])
		newY, _ := strconv.Atoi(newDirection[0])
		nextCoordinates = []int{nextCoordinates[0] + newY, nextCoordinates[1] + newX}
		direction = newDirection[2]
		count = count + 1
	}

	return count / 2
}

func FindConnectingPipe(matrix [][]string, point []int) ([]int, string) {
	newPoint := make([]int, 0)
	//Oben
	if point[0] > 0 {
		if matrix[point[0]-1][point[1]] == "|" {
			newPoint = append(newPoint, point[0]-1)
			newPoint = append(newPoint, point[1])
			return newPoint, "unten"
		}
		if matrix[point[0]-1][point[1]] == "F" {
			newPoint = append(newPoint, point[0]-1)
			newPoint = append(newPoint, point[1])
			return newPoint, "unten"
		}
		if matrix[point[0]-1][point[1]] == "7" {
			newPoint = append(newPoint, point[0]-1)
			newPoint = append(newPoint, point[1])
			return newPoint, "unten"
		}
	}
	//Unten
	if point[0] < len(matrix) {
		if matrix[point[0]+1][point[1]] == "|" {
			newPoint = append(newPoint, point[0]+1)
			newPoint = append(newPoint, point[1])
			return newPoint, "oben"
		}
		if matrix[point[0]+1][point[1]] == "L" {
			newPoint = append(newPoint, point[0]+1)
			newPoint = append(newPoint, point[1])
			return newPoint, "oben"
		}
		if matrix[point[0]+1][point[1]] == "J" {
			newPoint = append(newPoint, point[0]+1)
			newPoint = append(newPoint, point[1])
			return newPoint, "oben"
		}
	}
	//Links
	if point[1] > 0 {
		if matrix[point[0]][point[1]-1] == "-" {
			newPoint = append(newPoint, point[0])
			newPoint = append(newPoint, point[1]-1)
			return newPoint, "rechts"
		}
	}
	//Rechts
	if point[1] < len(matrix[point[0]]) {
		if matrix[point[0]][point[1]+1] == "-" {
			newPoint = append(newPoint, point[0])
			newPoint = append(newPoint, point[1]+1)
			return newPoint, "links"
		}
	}
	return newPoint, ""
}

func SplitLine(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func SplitLines(s []string) [][]string {
	matrix := make([][]string, len(s))
	for i, line := range s {
		for _, character := range line {
			matrix[i] = append(matrix[i], string(character))
		}
	}
	return matrix
}
