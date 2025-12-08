package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("cmd/06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	var questions [][]string

	for _, line := range lines {
		fields := strings.Fields(line)

		for i, field := range fields {
			if i >= len(questions) {
				questions = append(questions, []string{})
			}
			questions[i] = append(questions[i], field)
		}
	}

	cumulativeSum := 0

	for _, question := range questions {
		cumulativeSum += answerQuestion(question)
	}

	fmt.Println(cumulativeSum)
}

func answerQuestion(question []string) int {
	operator := question[len(question)-1]
	operands := question[1 : len(question)-1]

	answer, _ := strconv.Atoi(question[0])
	for _, operand := range operands {
		operandInt, _ := strconv.Atoi(operand)
		switch operator {
		case "+":
			answer += operandInt
		case "*":
			answer *= operandInt
		}
	}

	return answer
}
