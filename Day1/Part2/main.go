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

func readCalibrationNumbers(input []string) []string {

	var remember string
	sums := make([]string, 0)
	var firstDigit int
	var secondDigit int

	for i := 0; i < len(input); i++ {
		wholeLine := input[i]
		fmt.Println("WholeLine: ", wholeLine)
		for j := 0; j < len(wholeLine); j++ {
			remember = remember + string(wholeLine[j])
			firstDigit = checkString(remember)
			if firstDigit != -1 {
				remember = ""
				break
			}
		}
		for j := len(wholeLine) - 1; j >= 0; j-- {
			remember = remember + string(wholeLine[j])
			secondDigit = checkString(remember)
			if secondDigit != -1 {
				remember = ""
				break
			}
		}
		newPairForSums := strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit)
		fmt.Println(newPairForSums)
		sums = append(sums, newPairForSums)
	}
	return sums
}

func checkString(remember string) int {

	digits := map[string]int{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"orez":  0,
		"eno":   1,
		"owt":   2,
		"eerht": 3,
		"ruof":  4,
		"evif":  5,
		"xis":   6,
		"neves": 7,
		"thgie": 8,
		"enin":  9,
	}
	var returnDigit int

	for key, value := range digits {
		if strings.Contains(remember, key) {
			returnDigit = value
			fmt.Println(returnDigit)
			return returnDigit
		}
	}
	return -1
}

func calculateCalibrationSum(sums []string) int {

	var calibrationSum int
	for i := 0; i < len(sums); i++ {
		foo,_ := strconv.Atoi(sums[i])
		calibrationSum = calibrationSum + foo
	}
	return calibrationSum
}