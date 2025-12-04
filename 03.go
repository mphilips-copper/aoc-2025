package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("03.txt")
	if err != nil {
		log.Fatal(err)
	}

	batteryBanks := strings.Split(string(input), "\n")
	joltage := 0

	for _, batteryBank := range batteryBanks {
		lD12, index := findLargestDigitAndIndex(0, trimLastNChars(batteryBank, 11))
		lD11, index := findLargestDigitAndIndex(index+1, trimLastNChars(batteryBank, 10))
		lD10, index := findLargestDigitAndIndex(index+1, trimLastNChars(batteryBank, 9))
		lD9, index := findLargestDigitAndIndex(index+1, trimLastNChars(batteryBank, 8))
		lD8, index := findLargestDigitAndIndex(index+1, trimLastNChars(batteryBank, 7))
		lD7, index := findLargestDigitAndIndex(index+1, trimLastNChars(batteryBank, 6))
		lD6, index := findLargestDigitAndIndex(index+1, trimLastNChars(batteryBank, 5))
		lD5, index := findLargestDigitAndIndex(index+1, trimLastNChars(batteryBank, 4))
		lD4, index := findLargestDigitAndIndex(index+1, trimLastNChars(batteryBank, 3))
		lD3, index := findLargestDigitAndIndex(index+1, trimLastNChars(batteryBank, 2))
		lD2, index := findLargestDigitAndIndex(index+1, trimLastNChars(batteryBank, 1))
		lD1, _ := findLargestDigitAndIndex(index+1, batteryBank)

		// lol, lmao even
		joltage = joltage +
			lD12*100000000000 +
			lD11*10000000000 +
			lD10*1000000000 +
			lD9*100000000 +
			lD8*10000000 +
			lD7*1000000 +
			lD6*100000 +
			lD5*10000 +
			lD4*1000 +
			lD3*100 +
			lD2*10 +
			lD1
	}

	fmt.Println(joltage)
}

func trimLastNChars(s string, n int) string {
	return s[:len(s)-n]
}

func findLargestDigitAndIndex(startAt int, batteryBank string) (int, int) {
	largestDigit := 0
	index := 0

	for i, r := range batteryBank[startAt:] {
		// Convert rune to numeric digit value (e.g., '5' -> 5)
		digitValue := int(r - '0')
		if digitValue > largestDigit {
			largestDigit = digitValue
			index = i
		}
	}

	return largestDigit, startAt + index
}
