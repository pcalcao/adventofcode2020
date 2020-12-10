package main

import (
	"fmt"
	"log"
)

func searchRow(direction rune, lower int, upper int) (int, int) {
	var low int
	var high int
	if direction == 'F' || direction == 'L' {
		low = lower
		high = (upper-lower)/2 + lower
	} else {
		high = upper
		low = (upper-lower)/2 + lower + 1
	}
	return low, high
}

func breakTicket(ticket string) (string, string) {
	// ticket example: BFFFBBFRRR
	// Should be broken in BFFFBBF, RRR
	row := ticket[0:7]
	seat := ticket[7:len(ticket)]
	return row, seat
}

func getRow(rows string) int {
	lower, upper := 0, 127
	for _, char := range rows {
		lower, upper = searchRow(char, lower, upper)
	}
	return lower
}

func getSeat(seats string) int {
	lower, upper := 0, 7
	for _, char := range seats {
		lower, upper = searchRow(char, lower, upper)
	}
	return lower
}

func getSeatIDFromTicket(ticket string) int {
	rows, seats := breakTicket(ticket)
	row := getRow(rows)
	seat := getSeat(seats)

	return calculateSeatID(row, seat)
}

func calculateSeatID(row int, seat int) int {
	return row*8 + seat
}

func day5_1(lines []string) int {
	maxID := 0
	for _, line := range lines {
		id := getSeatIDFromTicket(line)
		if id > maxID {
			maxID = id
		}
	}
	return maxID
}

func day5_2(lines []string) int {
	set := createSet([]int{})
	var candidate int

	for _, line := range lines {
		id := getSeatIDFromTicket(line)
		set[id] = struct{}{}
	}
	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			maybeID := calculateSeatID(i, j)
			if _, ok := set[maybeID]; !ok {
				// check if before and after are in there
				_, right := set[maybeID+1]
				_, left := set[maybeID-1]
				if right && left {
					candidate = maybeID
				}
			}
		}
	}
	return candidate
}

func day5Main() {
	lines, err := ReadLines("../inputs/day5")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	count := day5_1(lines)
	fmt.Println(count)

	count2 := day5_2(lines)
	fmt.Println(count2)
}
