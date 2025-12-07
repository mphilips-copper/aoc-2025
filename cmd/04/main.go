package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type gridSlot struct {
	state  rune
	xCoord int
	yCoord int
}

type grid struct {
	width     int
	height    int
	gridSlots [][]gridSlot
}

const (
	EMPTY = '.'
	PAPER = '@'
)

func main() {
	input, err := os.ReadFile("cmd/04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	gridLines := strings.Split(string(input), "\n")
	paperGrid := grid{
		width:     len(gridLines[0]),
		height:    len(gridLines),
		gridSlots: make([][]gridSlot, len(gridLines)),
	}

	for i, line := range gridLines {
		paperGrid.gridSlots[i] = make([]gridSlot, paperGrid.width)

		for j, char := range line {
			paperGrid.gridSlots[i][j] = gridSlot{
				xCoord: j,
				yCoord: i,
				state:  char,
			}
		}
	}

	totalSlotsCleared := 0
	numSlotsCleared := 1 // start as some non-zero value
	for numSlotsCleared > 0 {
		paperAndForkliftableGridSlots := paperAndForkliftableGridSlots(paperGrid)

		numSlotsCleared = len(paperAndForkliftableGridSlots)
		totalSlotsCleared += numSlotsCleared

		for _, slot := range paperAndForkliftableGridSlots {
			paperGrid.gridSlots[slot.yCoord][slot.xCoord].state = EMPTY
		}
	}

	fmt.Println(totalSlotsCleared)
}

func paperAndForkliftableGridSlots(paperGrid grid) []gridSlot {
	paperAndForkliftableGridSlots := []gridSlot{}

	for i := 0; i < paperGrid.height; i++ {
		for j := 0; j < paperGrid.width; j++ {
			if paperGrid.gridSlots[i][j].state == PAPER && numEmptyNeighbors(paperGrid, paperGrid.gridSlots[i][j]) > 4 {
				paperAndForkliftableGridSlots = append(paperAndForkliftableGridSlots, paperGrid.gridSlots[i][j])
			}
		}
	}

	return paperAndForkliftableGridSlots
}

func numEmptyNeighbors(paperGrid grid, slot gridSlot) int {
	slotNeighbors := neighbors(paperGrid, slot)
	// 8 is the most possible neighbors,
	// any out-of-bounds neighbors did not return and we know they are empty
	numEmptyNeighbors := 8 - len(slotNeighbors)

	for _, neighbor := range slotNeighbors {
		if neighbor.state == EMPTY {
			numEmptyNeighbors++
		}
	}

	return numEmptyNeighbors
}

func neighbors(paperGrid grid, slot gridSlot) []gridSlot {
	neighbors := []gridSlot{}

	if slot.xCoord-1 >= 0 && slot.yCoord-1 >= 0 {
		neighbors = append(neighbors, paperGrid.gridSlots[slot.yCoord-1][slot.xCoord-1])
	}
	if slot.yCoord-1 >= 0 {
		neighbors = append(neighbors, paperGrid.gridSlots[slot.yCoord-1][slot.xCoord])
	}
	if slot.xCoord+1 < paperGrid.width && slot.yCoord-1 >= 0 {
		neighbors = append(neighbors, paperGrid.gridSlots[slot.yCoord-1][slot.xCoord+1])
	}
	if slot.xCoord-1 >= 0 {
		neighbors = append(neighbors, paperGrid.gridSlots[slot.yCoord][slot.xCoord-1])
	}
	// skip xCoord, yCoord - you aren't your own neighbor
	if slot.xCoord+1 < paperGrid.width {
		neighbors = append(neighbors, paperGrid.gridSlots[slot.yCoord][slot.xCoord+1])
	}
	if slot.xCoord-1 >= 0 && slot.yCoord+1 < paperGrid.height {
		neighbors = append(neighbors, paperGrid.gridSlots[slot.yCoord+1][slot.xCoord-1])
	}
	if slot.yCoord+1 < paperGrid.height {
		neighbors = append(neighbors, paperGrid.gridSlots[slot.yCoord+1][slot.xCoord])
	}
	if slot.xCoord+1 < paperGrid.width && slot.yCoord+1 < paperGrid.height {
		neighbors = append(neighbors, paperGrid.gridSlots[slot.yCoord+1][slot.xCoord+1])
	}

	return neighbors
}
