package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"unicode/utf8"
)

func main() {
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open("01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numZeroes := 0
	fmt.Println(numZeroes)

	dial := 50

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// assume perfect input - "R" prefix means travel positively...
		direction := 1

		// ...and "L" prefix means travel negatively
		if strings.HasPrefix(line, "L") {
			direction = -1
		}

		line = trimFirstRune(line)

		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		clicks := i * direction

		dial = turn(dial, clicks)
		fmt.Println(dial)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func turn(val int, clicks int) int {
	// don't care about the first
	// lmao actually we do
	clicks = clicks % 100
	val = val + clicks

	return val
}

// Source - https://stackoverflow.com/a/48801414
// Posted by user5728991, modified by community. See post 'Timeline' for change history
// Retrieved 2025-12-01, License - CC BY-SA 4.0

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
