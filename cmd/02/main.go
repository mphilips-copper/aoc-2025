package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	stringRanges := strings.Split(string(input), ",")
	sumInvalidIds := 0

	uniqueIds := make(map[int]bool)

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
				if uniqueIds[key] == false {
					uniqueIds[key] = true
					sumInvalidIds = sumInvalidIds + key
				}
			}
		}
	}

	fmt.Println(sumInvalidIds)
}

func containsRepeatedSequence(input string) bool {
	for subLength := range len(input) - 1 {
		if len(input)%(subLength+1) != 0 {
			continue
		}

		stringChunks := SplitSubN(input, subLength+1)

		if allChunksEqual(stringChunks) {
			return true
		}
	}

	return false
}

// Source - https://stackoverflow.com/a/39347212
// Posted by mozey
// Retrieved 2025-12-02, License - CC BY-SA 3.0
func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

func allChunksEqual(chunks []string) bool {
	comparisonChunk := chunks[0]

	for i := 1; i < len(chunks); i++ {
		if chunks[i] != comparisonChunk {
			return false
		}
	}

	return true
}

