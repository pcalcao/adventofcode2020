package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func isPasswordValid(min int, max int, letter string, pwd string) bool {
	occurences := strings.Count(pwd, letter)
	return occurences <= max && occurences >= min
}

func isPasswordValidV2(p1 int, p2 int, letter string, pwd string) bool {
	// Basically doing an XOR but Go is shit
	return (pwd[p1-1] == letter[0]) != (pwd[p2-1] == letter[0])
}

func breakStringEntry(entry string) (int, int, string, string) {
	re := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	matches := re.FindStringSubmatch(entry)

	minMatches, _ := strconv.Atoi(matches[1])
	maxMatches, _ := strconv.Atoi(matches[2])
	return minMatches, maxMatches, matches[3], matches[4]
}

func day2_1(lines []string) int {
	count := 0
	for _, entry := range lines {

		min, max, letter, pwd := breakStringEntry(entry)

		if isPasswordValid(min, max, letter, pwd) {
			count++
		}
	}
	return count
}

func day2_2(lines []string) int {
	count := 0
	for _, entry := range lines {

		p1, p2, letter, pwd := breakStringEntry(entry)

		if isPasswordValidV2(p1, p2, letter, pwd) {
			count++
		}
	}
	return count
}

func day2Main() {
	lines, err := ReadLines("../inputs/day2")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	count := day2_1(lines)
	fmt.Println(count)

	count2 := day2_2(lines)
	fmt.Println(count2)

	p1, p2, letter, pwd := breakStringEntry("2-9 c: ccccccccc")
	fmt.Println(isPasswordValidV2(p1, p2, letter, pwd))

}
