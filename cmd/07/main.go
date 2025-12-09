package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("cmd/07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Implement solution

	fmt.Println(string(input))
}
