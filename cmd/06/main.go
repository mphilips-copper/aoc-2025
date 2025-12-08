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

	questions := parseQuestions(input)

	cumulativeSum := 0

	for _, question := range questions {
		cumulativeSum += answerQuestion(question)
	}

	fmt.Println(cumulativeSum)
}

func parseQuestions(input []byte) [][]string {
	lines := strings.Split(string(input), "\n")

	if !(len(lines[0]) == len(lines[1]) &&
		len(lines[1]) == len(lines[2]) &&
		len(lines[2]) == len(lines[3]) &&
		len(lines[3]) == len(lines[4])) {
		log.Fatal("nah bro")
	}

	var questions [][]string

	idx := len(lines[0]) - 1
	questionsIdx := 0
	for idx >= 0 {
		// brittle, expects the input to be exactly 5 lines
		zero := string(lines[0][idx])
		one := string(lines[1][idx])
		two := string(lines[2][idx])
		three := string(lines[3][idx])
		operator := string(lines[4][idx])

		// the current question is 'over', start a new one
		if zero == " " && one == " " && two == " " && three == " " && operator == " " {
			questionsIdx += 1
			idx -= 1
			continue
		}

		// avoid nil slice
		if questionsIdx >= len(questions) {
			questions = append(questions, []string{})
		}

		// combine the digits into an operand
		operand := strings.TrimSpace(zero + one + two + three)
		questions[questionsIdx] = append(questions[questionsIdx], operand)

		// if there's an operator, add it to the question
		if operator != " " {
			questions[questionsIdx] = append(questions[questionsIdx], operator)
		}

		idx -= 1
	}

	return questions
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
