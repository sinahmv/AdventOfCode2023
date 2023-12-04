package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day4/input.txt")
	input := string(file)
	cards := SplitLines(input)
	fmt.Println("Ergebnis Part 1: ", part1(cards))
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func part1(cards []string) int {

	rememberWinningNumbers := make([]int, 0)
	rememberDrawingNumbers := make([]int, 0)
	score := -1
	endscore := 0

	for i := 0; i < len(cards); i++ {

		currentInput := cards[i]
		currentInput = strings.Replace(currentInput, "  ", " ", -1)

		splitInput := strings.Split(currentInput, ": ")
		splitNumbers := strings.Split(splitInput[1], " | ")
		splitWinningNumbers := strings.Split(splitNumbers[0], " ")
		splitDrawingNumbers := strings.Split(splitNumbers[1], " ")

		for i := 0; i < len(splitWinningNumbers); i++ {
			WinningNumber, _ := strconv.Atoi(splitWinningNumbers[i])
			rememberWinningNumbers = append(rememberWinningNumbers, WinningNumber)
		}

		for s := 0; s < len(splitDrawingNumbers); s++ {
			DrawingNumber, _ := strconv.Atoi(splitDrawingNumbers[s])
			rememberDrawingNumbers = append(rememberDrawingNumbers, DrawingNumber)
		}

		for j := 0; j < len(splitDrawingNumbers); j++ {
			currentNumber,_ := strconv.Atoi(splitDrawingNumbers[j])
			for k := 0; k < len(rememberWinningNumbers); k++ {
				numb := rememberWinningNumbers[k]
				if numb == currentNumber && numb != 0 {
					if score == -1 {
						score = 1
					} else {
						score = score * 2
					}
				}
			}
		}
		if score != -1 {
			endscore = endscore + score
		}
		rememberWinningNumbers = make([]int, 0)
		score = -1
	}
	return endscore
}
