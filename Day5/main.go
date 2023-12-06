package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day5/input.txt")
	input := string(file)
	lines := SplitLines(input)
	part1(lines)
	//fmt.Println("Ergebnis Part 2: ", part1(lines))
}

func part1(lines []string) {
	getMySeeds(lines)
	getMyMaps(lines)
}

func getMySeeds(lines []string) []int {
	seedNumbers := make([]int, 0)
	splitInput := strings.Split(lines[0], ":")
	seedsAsString := splitInput[1]
	seeds := strings.Fields(seedsAsString)
	for i := 0; i < len(seeds); i++ {
		currentSeed, _ := strconv.Atoi(seeds[i])
		seedNumbers = append(seedNumbers, currentSeed)
	}
	fmt.Println("SeedNumbers: ", seedNumbers)
	return seedNumbers
}

func getMyMaps(lines []string) {

	for _, line := range lines {
		//Parsing Input
		inputLine := strings.Fields(line)

		fmt.Println(inputLine)

		for i,numb := range inputLine {
			if numb != 
		}

		for _, seed := range getMySeeds(lines) {
			fmt.Println("Seed: ", seed)
			for _, numb := range inputLine {

			}
		}
	}
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
