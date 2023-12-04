package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Liefert aus dem Input File den Inhalt als []byte
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day2/input.txt")
	//Input an WhiteSpaces splitten
	input := string(file)
	games := SplitLines(input)
	fmt.Println("Ergebnis Part 1: ", sumUp(part1(games)))
	fmt.Println("Ergebnis Part 2: ", sumUp(part2(games)))
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func part1(games []string) []int {

	numbersOfPossibleGames := make([]int, 0)
	var gameID int

	for i := 0; i < len(games); i++ {

		var game string
		winGame := true

		currentInput := games[i]

		splitInput := strings.Split(currentInput, " ")
		game, _, _ = strings.Cut(splitInput[1], ":")

		gameID, _ = strconv.Atoi(game)

		for j := 0; j < len(splitInput); j++ {
			currentString := splitInput[j]

			if strings.Contains(currentString, "blue") {
				count, _ := strconv.Atoi(splitInput[j-1])
				if count > 14 {
					winGame = false
				}
			} else if strings.Contains(currentString, "red") {
				count, _ := strconv.Atoi(splitInput[j-1])
				if count > 12 {
					winGame = false
				}
			} else if strings.Contains(currentString, "green") {
				count, _ := strconv.Atoi(splitInput[j-1])
				if count > 13 {
					winGame = false
				}
			}
		}
		if winGame {
			numbersOfPossibleGames = append(numbersOfPossibleGames, gameID)
		}
	}
	return numbersOfPossibleGames
}

func part2(games []string) []int {

	powers := make([]int, 0)
	multiplyCounts := 0

	for i := 0; i < len(games); i++ {

		currentLowestRed := -1
		currentLowestBlue := -1
		currentLowestGreen := -1

		currentInput := games[i]

		splitInput := strings.Split(currentInput, " ")

		for j := 0; j < len(splitInput); j++ {
			currentString := splitInput[j]

			if strings.Contains(currentString, "blue") {
				count, _ := strconv.Atoi(splitInput[j-1])
				if count > currentLowestBlue {
					currentLowestBlue = count
				}
			} else if strings.Contains(currentString, "red") {
				count, _ := strconv.Atoi(splitInput[j-1])
				if count > currentLowestRed {
					currentLowestRed = count
				}
			} else if strings.Contains(currentString, "green") {
				count, _ := strconv.Atoi(splitInput[j-1])
				if count > currentLowestGreen {
					currentLowestGreen = count
				}
			}
		}
		multiplyCounts = currentLowestBlue * currentLowestGreen * currentLowestRed
		powers = append(powers, multiplyCounts)
	}
	return powers
}

func checkIfGameIsPossible(countRed int, countBlue int, countGreen int) bool {
	gameIsPossible := false
	if countBlue <= 14 && countGreen <= 13 && countRed <= 12 {
		gameIsPossible = true
	}
	return gameIsPossible
}

func sumUp(numbersOfPossibleGames []int) int {
	sum := 0
	for i := 0; i < len(numbersOfPossibleGames); i++ {
		sum = sum + numbersOfPossibleGames[i]
	}
	return sum
}
