package main

import (
	"fmt"
	"log"
)

func readAnswers(lines []string) []map[rune]struct{} {
	var answers []map[rune]struct{}

	currentTicket := make(map[rune]struct{})
	for _, line := range lines {
		if line == "" {
			answers = append(answers, currentTicket)
			currentTicket = make(map[rune]struct{})
		}

		for _, c := range line {
			currentTicket[c] = struct{}{}
		}
	}
	answers = append(answers, currentTicket)
	return answers
}

func readAnswersCommon(lines []string) int {
	var answers []map[rune]int

	currentTicket := make(map[rune]int)
	numLines := 0
	count := 0
	for _, line := range lines {
		if line == "" {
			// filter out all the answers that don't match numLines
			for _, number := range currentTicket {
				if number == numLines {
					count++
				}
			}
			answers = append(answers, currentTicket)
			currentTicket = make(map[rune]int)
			numLines = 0
		} else {
			numLines++
		}
		for _, c := range line {
			answersSoFar, ok := currentTicket[c]
			if !ok {
				currentTicket[c] = 1
			}
			currentTicket[c] = answersSoFar + 1
		}
		fmt.Println(numLines, currentTicket)
	}
	for _, number := range currentTicket {
		if number == numLines {
			count++
		}
	}
	answers = append(answers, currentTicket)
	fmt.Println(answers)
	return count
}

func sumAnswers(answers []map[rune]struct{}) int {
	count := 0
	for _, answer := range answers {
		count += len(answer)
	}
	return count
}

func day6_1(lines []string) int {
	return sumAnswers(readAnswers(lines))
}

func day6_2(lines []string) int {
	return readAnswersCommon(lines)
}

func day6Main() {
	lines, err := ReadLines("../inputs/day6")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	count := day6_1(lines)
	fmt.Println(count)

	count2 := day6_2(lines)
	fmt.Println(count2)
}
