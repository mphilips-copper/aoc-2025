package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type event struct {
	id       int
	coverage int
}

func main() {
	input, err := os.ReadFile("cmd/05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	freshIDRanges, _ := parseInput(lines)

	// An algorithm to build a list of covered ranges from a given set of
	// intervals, effectively merging overlapping or adjacent intervals into a
	// consolidated list of non-overlapping covered ranges, can be implemented
	// using a sweep-line approach.

	events := []event{}

	for _, freshIDRange := range freshIDRanges {
		events = append(events, event{
			id:       freshIDRange[0],
			coverage: 1,
		})
		events = append(events, event{
			id:       freshIDRange[1],
			coverage: -1,
		})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].id == events[j].id {
			return events[i].coverage > events[j].coverage
		}
		return events[i].id < events[j].id
	})

	coverageCount := 0
	currentStart := 0
	coveredRanges := [][]int{}

	for _, event := range events {
		if coverageCount == 0 && event.coverage == 1 {
			currentStart = event.id
		}
		coverageCount += event.coverage
		if coverageCount == 0 && event.coverage == -1 {
			coveredRanges = append(coveredRanges, []int{currentStart, event.id})
		}
	}

	coveredCount := 0
	for _, coveredRange := range coveredRanges {
		coveredCount += coveredRange[1] - coveredRange[0] + 1
	}
	fmt.Println(coveredCount)
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
