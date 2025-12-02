package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
)

func main() {
	// https://pkg.go.dev/container/list
	l := list.New()
	for i := range 101 {
		l.PushBack(i)
	}

	// To achieve circular traversal, you need to handle the end conditions
	// explicitly. When traversing forward, if e.Next() becomes nil, you can reset
	// e to l.Front(). Similarly, for backward traversal, if e.Prev() becomes nil, you can reset e to l.Back().

	arrow := l.Front()
	fmt.Println(arrow.Value)
	arrow.turnRight(50)
	fmt.Println(arrow.Value)

	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open("01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	numZeroes := 0
	fmt.Println(numZeroes)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (e *Element) turnRight(clicks int) Element {
	for i := range clicks {
		e = e.Next()
	}

	return e
}

// func (e *Element) turnLeft(clicks int) Element {

// }
