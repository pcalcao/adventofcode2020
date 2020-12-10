package main

import (
	"fmt"
	"log"
)

func day7_1(lines []string) int {
	return 0
}

func day7_2(lines []string) int {
	return 0
}

func day7Main() {
	lines, err := ReadLines("../inputs/day7")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	count := day7_1(lines)
	fmt.Println(count)

	count2 := day7_2(lines)
	fmt.Println(count2)
}
