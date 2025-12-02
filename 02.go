package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("02.txt")
	if err != nil {
		log.Fatal(err)
	}

	stringRanges := strings.Split(string(input), ",")
	sumInvalidIds := 0

	for _, stringRange := range stringRanges {
		stringIds := strings.Split(stringRange, "-")
		// let's assume input is perfect and there will be exactly 2 ids here
		start, _ := strconv.Atoi(stringIds[0])
		end, _ := strconv.Atoi(stringIds[1])

		ids := make(map[int]string)
		for ; start <= end; start++ {
			ids[start] = strconv.Itoa(start)
		}

		for key, value := range ids {
			if containsRepeatedSequence(value) {
				fmt.Println("invalid key:", key)
				sumInvalidIds = sumInvalidIds + key
			}
		}

		fmt.Println("\n")
	}

	fmt.Println(sumInvalidIds)
}

func containsRepeatedSequence(input string) bool {
	// fmt.Println("input:", input)

	// digit frequency cases
	// only one digit - TRUE
	// two+ digits - TRUE if all counts >=2 and all counts ==

	digitFrequency := make(map[string]int)
	for _, char := range input {
		digitChar := string(char)
		digitFrequency[digitChar]++
	}

	// fmt.Println(digitFrequency)

	if len(digitFrequency) == 1 {
		return true
	}

	firstCount := digitFrequency[string(input[0])]
	// fmt.Println("first char:", string(input[0]), "first count:", firstCount)

	for _, value := range digitFrequency {
		if value < 2 || value != firstCount {
			return false
		}
	}

	return true
}
