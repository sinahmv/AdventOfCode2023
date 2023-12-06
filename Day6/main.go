package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day6/input.txt")
	input := string(file)
	lines := SplitLines(input)
	fmt.Println("Ergebnis Part 1: ", part1(lines))
	fmt.Println("Ergebnis Part 2: ", part2(lines))
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func part1(input []string) int {

	raceTimes := make([]string, 0)
	highsores := make([]string, 0)
	output := make([]int, 0)
	waysToWin := 0
	final := 1

	for _, line := range input {

		splitInput := strings.Fields(line)

		for _, item := range splitInput {

			if strings.Contains(line, "Time:") {
				raceTimes = append(raceTimes, string(item))
			} else {
				highsores = append(highsores, string(item))
			}
		}
	}

	for i := 1; i < len(raceTimes); i++ {

		currentRaceTime, _ := strconv.Atoi(raceTimes[i])
		oldHighScore, _ := strconv.Atoi(highsores[i])

		for j := 0; j <= currentRaceTime; j++ {

			//Wie wird der Score berechnet???
			score := (currentRaceTime - j) * j

			if score > oldHighScore {
				waysToWin = waysToWin + 1
			}
		}
		output = append(output, waysToWin)
		waysToWin = 0
	}
	for _, item := range output {
		final = final * item
	}
	return final
}

func part2(input []string) int {

	raceTimes := make([]string, 0)
	highsores := make([]string, 0)
	raceTimeString := ""
	highsoreString := ""
	waysToWin := 0

	for _, line := range input {

		splitInput := strings.Fields(line)

		for _, item := range splitInput {

			if strings.Contains(line, "Time:") {
				raceTimes = append(raceTimes, string(item))
			} else {
				highsores = append(highsores, string(item))
			}
		}
	}

	for i := 1; i < len(raceTimes); i++ {
		raceTimeString = raceTimeString + raceTimes[i]
	}
	raceTime, _ := strconv.Atoi(raceTimeString)

	for i := 1; i < len(highsores); i++ {
		highsoreString = highsoreString + highsores[i]
	}
	highsore, _ := strconv.Atoi(highsoreString)

	fmt.Println(raceTime)
	fmt.Println(highsore)

	for j := 0; j <= raceTime; j++ {

		//Wie wird der Score berechnet???
		score := (raceTime - j) * j

		if score > highsore {
			waysToWin = waysToWin + 1
		}
	}
	return waysToWin
}
