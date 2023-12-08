package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

func main() {
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day7/input.txt")
	input := string(file)
	lines := SplitLines(input)

	//Input as Hands
	hands := make([]Hand, 0)

	for _, line := range lines {
		currentLine := strings.Fields(line)
		currentCards := currentLine[0]
		bid, _ := strconv.Atoi(currentLine[1])
		var hand Hand
		hand.cards = currentCards
		hand.bid = bid
		hands = append(hands, hand)
	}
	fmt.Println("Ergebnis Part 1:", bubbleSort(hands))
}

func bubbleSort(hands []Hand) int {

	swapped := true
	output := 0

	for swapped {
		swapped = false

		for i := 1; i < len(hands); i++ {
			if compareHands(hands[i], hands[i-1]) {
				hands[i], hands[i-1] = hands[i-1], hands[i]
				swapped = true
			}
		}
	}

	for s := 1; s <= len(hands); s++ {
		bid := hands[s-1].bid
		fmt.Println("Cards: ", hands[s-1].cards)
		output = output + (bid * s)
	}
	return output
}

func compareHands(hand1 Hand, hand2 Hand) bool {

	digits := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}
	hand1Score := getTypeOfHand(hand1)
	hand2Score := getTypeOfHand(hand2)

	if hand1Score < hand2Score {
		return true
	} else if hand1Score > hand2Score {
		return false
	} else if hand1Score == hand2Score {
		for i := range hand1.cards {
			if digits[string(hand1.cards[i])] < digits[string(hand2.cards[i])] {
				return true
			} else if digits[string(hand1.cards[i])] > digits[string(hand2.cards[i])] {
				return false
			} else {
				continue
			}
		}
	}
	return false
}

func getTypeOfHand(hand Hand) int {

	amountOfCards := make(map[string]int)

	for i := 0; i < len(hand.cards); i++ {

		cardValue := string(hand.cards[i])
		if _, exists := amountOfCards[cardValue]; exists {
			amountOfCards[cardValue] = amountOfCards[cardValue] + 1
		} else {
			amountOfCards[cardValue] = 1
		}
	}

	if containsValue(amountOfCards, 5) {
		return 6
	} else if containsValue(amountOfCards, 4) {
		return 5
	} else if containsValue(amountOfCards, 3) && containsValue(amountOfCards, 2) {
		return 4
	} else if containsValue(amountOfCards, 3) {
		return 3
	} else if countOccurrences(amountOfCards, 2) == 2 {
		return 2
	} else if containsValue(amountOfCards, 2) {
		return 1
	} else {
		return 0
	}
}

func containsValue(m map[string]int, value int) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

func countOccurrences(m map[string]int, value int) int {
	count := 0
	for _, v := range m {
		if v == value {
			count++
		}
	}
	return count
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
