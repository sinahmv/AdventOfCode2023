package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day11/input.txt")
	input := string(file)
	lines := SplitLine(input)
	universe := SplitLines(lines)
	fmt.Println("Ergebnis Part 1: ", Part1(universe))
}

func Part1(universe [][]string) int {

	output := 0
	coordinatesOfGalaxies := make([][]int, 0)
	universe = insertLines(universe)
	universe = insertColumns(universe)

	//getting coordinates
	for i := 0; i < len(universe); i++ {
		for j := 0; j < len(universe[i]); j++ {
			if strings.Contains(universe[i][j], "#") {
				nextCoordinates := []int{i, j}
				coordinatesOfGalaxies = append(coordinatesOfGalaxies, nextCoordinates)
			}
		}
	}

	//compare coordinates to get shortest ways
	for i, galaxy1 := range coordinatesOfGalaxies {
		for _, galaxy2 := range coordinatesOfGalaxies[i+1:] {
			shortestWay := Abs(galaxy2[0]-galaxy1[0]) + Abs(galaxy2[1]-galaxy1[1])
			output = output + shortestWay
		}
	}

	return output
}

func insertLines(universe [][]string) [][]string {

	clearLines := make([]int, 0)
	lineContainsGalaxy := false
	shift := 0

	for i := 0; i < len(universe); i++ {
		for j := 0; j < len(universe[i]); j++ {
			if strings.Contains(universe[i][j], "#") {
				lineContainsGalaxy = true
			}
		}
		if lineContainsGalaxy == false {
			clearLines = append(clearLines, i)
		}
		lineContainsGalaxy = false
	}

	for _, lineNumb := range clearLines {
		universe = append(universe[:lineNumb+shift+1], universe[lineNumb+shift:]...)
		shift += 1
	}
	return universe
}

func insertColumns(universe [][]string) [][]string {

	clearColumns := make([]int, 0)
	columnContainsGalaxy := false
	numRows := len(universe)
	numCols := len(universe[0])
	shift := 0

	for col := 0; col < numCols; col++ {
		for row := 0; row < numRows; row++ {
			if strings.Contains(universe[row][col], "#") {
				columnContainsGalaxy = true
			}
		}
		if columnContainsGalaxy == false {
			clearColumns = append(clearColumns, col)
		}
		columnContainsGalaxy = false
	}

	for _, columnNumb := range clearColumns {
		for i := range universe {
			universe[i] = append(universe[i][:columnNumb+shift+1], universe[i][columnNumb+shift:]...)
		}
		shift = shift + 1
	}
	return universe
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
