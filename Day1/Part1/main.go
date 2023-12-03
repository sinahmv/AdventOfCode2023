package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Liefert aus dem Input File den Inhalt als []byte
	file, _ := os.ReadFile("/Users/sinah/Code/AdventOfCode2023/Day1/input.txt")
	//Input an Zeilenumbrüchen splitten
	input := strings.Fields(string(file))
	calibrationSums := readCalibrationNumbers(input)
	fmt.Println("CalibrationSum: ", calculateCalibrationSum(calibrationSums))
}

func readCalibrationNumbers(input []string) []int {

	sums := make([]int, 0)
	var firstDigit string
	var secondDigit string

	for i := 0; i < len(input); i++ {
		wholeLine := input[i]
		fmt.Println("WholeLine: ", wholeLine)
		for j := 0; j < len(wholeLine); j++ {
			current := string(wholeLine[j])
			if current == "0" || current == "1" || current == "2" || current == "3" || current == "4" || current == "5" || current == "6" || current == "7" || current == "8" || current == "9" {
				firstDigit = current
				fmt.Println("Current: ", firstDigit)
				break
			}
		}
		for k := len(wholeLine) - 1; k >= 0; k-- {
			current := string(wholeLine[k])
			if current == "0" || current == "1" || current == "2" || current == "3" || current == "4" || current == "5" || current == "6" || current == "7" || current == "8" || current == "9" {
				secondDigit = current
				fmt.Println("Current: ", secondDigit)
				break
			}
		}
		newPairForSums := firstDigit + secondDigit
		poop,_ := strconv.Atoi(newPairForSums)
		fmt.Println(newPairForSums)
		sums = append(sums, poop)
	}
	return sums
}

func calculateCalibrationSum(sums []int) int {
	var calibrationSum int
	for i := 0; i < len(sums); i++ {
		calibrationSum = calibrationSum + sums[i]
	}
	return calibrationSum
}
