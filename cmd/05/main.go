package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("cmd/05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	freshIDRanges, ingredientIDs := parseInput(lines)

	numFreshIngredients := 0
	for _, ingredientID := range ingredientIDs {
		for _, freshIDRange := range freshIDRanges {
			if freshIDRange[0] <= ingredientID && freshIDRange[1] >= ingredientID {
				numFreshIngredients++
				break
			}
		}
	}

	fmt.Println(numFreshIngredients)
}

func parseInput(lines []string) ([][]int, []int) {
	var freshIDRanges [][]int
	var ingredientIDs []int
	inFirstChunk := true

	for _, line := range lines {
		if line == "" {
			inFirstChunk = false
			continue
		}

		if inFirstChunk {
			parts := strings.Split(line, "-")

			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			freshIDRanges = append(freshIDRanges, []int{start, end})
		} else {
			ingredientID, _ := strconv.Atoi(line)
			ingredientIDs = append(ingredientIDs, ingredientID)
		}
	}

	return freshIDRanges, ingredientIDs
}
