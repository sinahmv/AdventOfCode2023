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
	fmt.Println("Ergebnis Part 2: ", Part2(lines))
}

func Part2(lines []string) int {

	connections := 0
	numbersOfConnectedParts := make([]int, 0)

	for i := 0; i < len(lines); i++ {

		currentLine := lines[i]

		for pos, char := range currentLine {

			if char == '*' {

				fmt.Println("YAY")

				/* 				compareDown, downNumber := CompareDown(lines, pos, i)
				   				fmt.Println(compareDown)
				   				compareUp, upNumber := CompareUp(lines, pos, i)
				   				fmt.Println(compareUp) */
				compareLeft, leftNumber := CompareLeft(currentLine, pos)
				fmt.Println(compareLeft)
				compareRight, rightNumber := CompareRight(currentLine, pos)
				fmt.Println(compareRight)

				fmt.Println("YAY")

				if compareLeft {
					numbersOfConnectedParts = append(numbersOfConnectedParts, leftNumber)
					connections = connections + 1
				} else if compareRight {
					numbersOfConnectedParts = append(numbersOfConnectedParts, rightNumber)
					connections = connections + 1
				} /* else if compareUp && connections < 2 {
					connections = connections + 1
					numbersOfConnectedParts = append(numbersOfConnectedParts, upNumber)
				} else if compareDown && connections < 2 {
					connections = connections + 1
					numbersOfConnectedParts = append(numbersOfConnectedParts, downNumber)
				} */
				connections = 0
			}
		}
	}
	fmt.Println("numbersOfConnectedParts: ", numbersOfConnectedParts)
	sum := 0
	for i := 0; i < len(numbersOfConnectedParts)-1; i++ {
		sum = sum + (numbersOfConnectedParts[i] * numbersOfConnectedParts[i+1])
	}
	return sum
}

func CompareLeft(currentLine string, pos int) (bool, int) {

	wholeNumberOfPart := ""
	isConnected := false

	for i := range currentLine {
		for pos > i {
			charLeft := currentLine[pos-i]
			charLeftString := string(currentLine[pos-i])
			if unicode.IsNumber(rune(charLeft)) {
				isConnected = true
				wholeNumberOfPart = wholeNumberOfPart + charLeftString
			} else {
				break
			}
		}
	}
	wholeNumberOfPartInt, _ := strconv.Atoi(wholeNumberOfPart)
	return isConnected, wholeNumberOfPartInt
}

func CompareRight(currentLine string, pos int) (bool, int) {

	wholeNumberOfPart := ""
	isConnected := false

	for i := range currentLine {
		if pos < len(currentLine)-1 {
			charRight := currentLine[pos+i]
			charRightString := string(currentLine[pos+i])
			if unicode.IsNumber(rune(charRight)) {
				isConnected = true
				wholeNumberOfPart = wholeNumberOfPart + charRightString
			} else {
				break
			}
		}
	}
	wholeNumberOfPartInt, _ := strconv.Atoi(wholeNumberOfPart)
	return isConnected, wholeNumberOfPartInt
}

/* func CompareUp(lines []string, pos int, line int) (bool, int) {

	wholeNumberOfPart := ""
	isConnected := false

	if line > 0 {
		lineAbove := lines[line-1]
		isConnectedLeft, numberUpLeft := CompareLeft(lineAbove, pos)
		isConnectedRight, numberUpRight := CompareRight(lineAbove, pos)
		if isConnectedLeft || isConnectedRight {
			isConnected = true
			wholeNumberOfPart = fmt.Sprint(numberUpLeft) + fmt.Sprint(numberUpRight)
		}
	}
	wholeNumberOfPartInt, _ := strconv.Atoi(wholeNumberOfPart)
	return isConnected, wholeNumberOfPartInt
} */

/* func CompareDown(lines []string, pos int, line int) (bool, int) {

	wholeNumberOfPart := ""
	isConnected := false

	if line < len(lines) {
		lineBelow := lines[line+1]
		isConnectedLeft, numberUpLeft := CompareLeft(lineBelow, pos)
		isConnectedRight, numberUpRight := CompareRight(lineBelow, pos)
		if isConnectedLeft || isConnectedRight {
			isConnected = true
			wholeNumberOfPart = fmt.Sprint(numberUpLeft) + fmt.Sprint(numberUpRight)
		}
	}
	wholeNumberOfPartInt, _ := strconv.Atoi(wholeNumberOfPart)
	return isConnected, wholeNumberOfPartInt
} */

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
