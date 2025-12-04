package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	input, err := os.ReadFile("03.txt")
	if err != nil {
		log.Fatal(err)
	}

	batteryBanks := strings.Split(string(input), "\n")
	joltage := 0

	for _, batteryBank := range batteryBanks {
		// can never use the last rune as the first digit
		trimmedBatteryBank := trimLastRune(batteryBank)

		// find the largest digit from trimmed battery bank
		largestDigitTens, index := findLargestDigitAndIndex(0, trimmedBatteryBank)

		// find the largest digit to its right in battery bank
		largestDigitOnes, _ := findLargestDigitAndIndex(index+1, batteryBank)

		joltage = joltage + largestDigitTens*10 + largestDigitOnes
	}

	fmt.Println(joltage)
}

func trimLastRune(s string) string {
	_, i := utf8.DecodeLastRuneInString(s)
	return s[:len(s)-i]
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

	return largestDigit, index
}
