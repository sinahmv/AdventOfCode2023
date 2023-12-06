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
	fmt.Println("Ergebnis Part 2: ", part2(cards))
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func part2(cards []string) int {

	//finalscore := 0
	//amount := len(cards)
	//amountofcards := make([]string, amount)

	for i, card := range cards {

		//Parsing Input
		splitInput := strings.Split(card, ":")
		splitNumbers := strings.Split(splitInput[1], "|")
		splitWinningNumbers := strings.Fields(splitNumbers[0])
		splitDrawingNumbers := strings.Fields(splitNumbers[1])

		scoreOfCard := 0

		for _, number := range splitDrawingNumbers {
			currentNumber, _ := strconv.Atoi(number)
			for _, numb := range splitWinningNumbers {
				winningNumber, _ := strconv.Atoi(numb)
				if winningNumber == currentNumber {
					scoreOfCard = scoreOfCard + 1
				}
			}
		}
		for s := 1; s <= scoreOfCard; s++ {
			cards = append(cards, "")
			insert(cards, i+s, cards[i+s])
			fmt.Println("AddedCard: ", cards[i+s])
		}
		fmt.Println("Cards: ", len(cards))
		scoreOfCard = 0
	}
	return len(cards)
}

func insert(a []string, index int, value string) []string {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
