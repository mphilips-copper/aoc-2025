package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	// https://pkg.go.dev/container/list
	l := list.New()
	for i := range 100 {
		l.PushBack(i)
	}

	arrow := l.Front()
	arrow, _ = turnRight(l, arrow, 50)
	numZeroes := 0
	newZeroes := 0

	fmt.Println("dial is:", arrow.Value)
	fmt.Println("new zeroes:", newZeroes)
	fmt.Println("num zeroes:", numZeroes)
	fmt.Println()

	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open("01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "R") {
			line = trimFirstRune(line)

			i, err := strconv.Atoi(line)
			fmt.Println("dir R:", i)
			if err != nil {
				panic(err)
			}

			arrow, newZeroes = turnRight(l, arrow, i)
			numZeroes = numZeroes + newZeroes

			fmt.Println("dial is:", arrow.Value)
			fmt.Println("new zeroes:", newZeroes)
			fmt.Println("num zeroes:", numZeroes)
			fmt.Println()
		} else if strings.HasPrefix(line, "L") {
			// shouldn't do this in both branches but w/e
			line = trimFirstRune(line)

			i, err := strconv.Atoi(line)
			fmt.Println("dir L:", i)
			if err != nil {
				panic(err)
			}

			arrow, newZeroes = turnLeft(l, arrow, i)
			numZeroes = numZeroes + newZeroes

			fmt.Println("dial is:", arrow.Value)
			fmt.Println("new zeroes:", newZeroes)
			fmt.Println("num zeroes:", numZeroes)
			fmt.Println()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("num zeroes:", numZeroes)
}

func turnRight(l *list.List, e *list.Element, clicks int) (*list.Element, int) {
	numZeroes := 0

	for _ = range clicks {
		if e.Next() != nil {
			e = e.Next()
		} else {
			e = l.Front()

			numZeroes++
		}
	}

	return e, numZeroes
}

func turnLeft(l *list.List, e *list.Element, clicks int) (*list.Element, int) {
	numZeroes := 0

	for _ = range clicks {
		if e.Prev() != nil {
			e = e.Prev()
			if e.Value == 0 {
				numZeroes++
			}
		} else {
			e = l.Back()
		}
	}

	return e, numZeroes
}

// Source - https://stackoverflow.com/a/48801414
// Posted by user5728991, modified by community. See post 'Timeline' for change history
// Retrieved 2025-12-01, License - CC BY-SA 4.0
func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
