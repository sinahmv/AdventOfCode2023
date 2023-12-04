package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day3/input.txt")
	input := string(file)
	lines := SplitLines(input)
	fmt.Println("Ergebnis Part 1: ", SearchNumber(lines))
	fmt.Println("Ergebnis Part 2: ")
}

func SearchNumber(lines []string) int {

	hasAConnection := false
	numbersOfConnectedParts := make([]int, 0)

	for i := 0; i < len(lines); i++ {

		currentLine := lines[i]
		wholeNumberOfPart := ""

		for pos, char := range currentLine {

			if unicode.IsDigit(char) {

				currentNumber := string(char)
				wholeNumberOfPart = wholeNumberOfPart + currentNumber

				if CompareDown(lines, pos, i) || CompareUp(lines, pos, i) || CompareRight(currentLine, pos) ||
					CompareLeft(currentLine, pos) || CompareDiagonalUpLeft(lines, pos, i, currentLine) ||
					CompareDiagonalUpRight(lines, pos, i, currentLine) ||
					CompareDiagonalDownLeft(lines, pos, i, currentLine) ||
					CompareDiagonalDownRight(lines, pos, i, currentLine) {
					hasAConnection = true
				}

				if pos == len(currentLine)-1 {

					if hasAConnection {
						numb, _ := strconv.Atoi(wholeNumberOfPart)
						numbersOfConnectedParts = append(numbersOfConnectedParts, numb)
					}
					wholeNumberOfPart = ""
					hasAConnection = false
				}

			} else if !unicode.IsDigit(char) {

				if hasAConnection {
					numb, _ := strconv.Atoi(wholeNumberOfPart)
					numbersOfConnectedParts = append(numbersOfConnectedParts, numb)
				}
				wholeNumberOfPart = ""
				hasAConnection = false
			}
		}
	}
	fmt.Println("numbersOfConnectedParts: ", numbersOfConnectedParts)
	sum := 0
	for i := 0; i < len(numbersOfConnectedParts); i++ {
		sum = sum + numbersOfConnectedParts[i]
	}
	return sum
}

func CompareLeft(currentLine string, pos int) bool {

	isConnected := false

	if pos > 0 {
		charLeftString := string(currentLine[pos-1])
		charLeft := currentLine[pos-1]
		if charLeftString != "." && !unicode.IsNumber(rune(charLeft)) {
			isConnected = true
		}
	}
	return isConnected
}

func CompareRight(currentLine string, pos int) bool {

	isConnected := false

	if pos < len(currentLine)-1 {
		charRightString := string(currentLine[pos+1])
		charRight := currentLine[pos+1]
		if charRightString != "." && !unicode.IsNumber(rune(charRight)) {
			isConnected = true
		}
	}
	return isConnected
}

func CompareUp(lines []string, pos int, line int) bool {

	isConnected := false

	if line > 0 {
		lineAbove := lines[line-1]
		charUpperString := string(lineAbove[pos])
		charUp := lineAbove[pos]
		if charUpperString != "." && !unicode.IsNumber(rune(charUp)) {
			isConnected = true
		}
	}
	return isConnected
}

func CompareDown(lines []string, pos int, line int) bool {

	isConnected := false

	if line < len(lines)-1 {
		lineBelow := lines[line+1]
		charBelowString := string(lineBelow[pos])
		charBelow := lineBelow[pos]
		if charBelowString != "." && !unicode.IsNumber(rune(charBelow)) {
			isConnected = true
		}
	}
	return isConnected
}

func CompareDiagonalUpLeft(lines []string, pos int, line int, currentLine string) bool {

	isConnected := false

	if line > 0 {
		lineAbove := lines[line-1]
		if pos > 0 {
			charBelowStringLeft := string(lineAbove[pos-1])
			charBelowLeft := lineAbove[pos-1]
			if charBelowStringLeft != "." && !unicode.IsNumber(rune(charBelowLeft)) {
				isConnected = true
			}
		}
	}
	return isConnected
}

func CompareDiagonalUpRight(lines []string, pos int, line int, currentLine string) bool {

	isConnected := false

	if line > 0 {
		lineAbove := lines[line-1]
		if pos < len(currentLine)-1 {
			charBelowStringRight := string(lineAbove[pos+1])
			charBelowRight := lineAbove[pos+1]
			if charBelowStringRight != "." && !unicode.IsNumber(rune(charBelowRight)) {
				isConnected = true
			}
		}
	}
	return isConnected
}

func CompareDiagonalDownRight(lines []string, pos int, line int, currentLine string) bool {

	isConnected := false

	if line < len(lines)-1 {
		lineBelow := lines[line+1]
		if pos < len(currentLine)-1 {
			charBelowStringRight := string(lineBelow[pos+1])
			charBelowRight := lineBelow[pos+1]
			if charBelowStringRight != "." && !unicode.IsNumber(rune(charBelowRight)) {
				isConnected = true
			}
		}
	}
	return isConnected
}

func CompareDiagonalDownLeft(lines []string, pos int, line int, currentLine string) bool {

	isConnected := false

	if line < len(lines)-1 {
		lineBelow := lines[line+1]
		if pos > 0 {
			charBelowStringLeft := string(lineBelow[pos-1])
			charBelowLeft := lineBelow[pos-1]
			if charBelowStringLeft != "." && !unicode.IsNumber(rune(charBelowLeft)) {
				isConnected = true
			}
		}
	}
	return isConnected
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
