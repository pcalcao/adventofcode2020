package main

import "log"
import "fmt"

func getNextPosition(row int, column int, length int, down int, right int) (int, int) {
	/*
	 * Moves the current position 1 down, 3 right
	 * wrapping around as needed
	 */
	newRow, newColumn := row+down, (column+right)%length

	return newRow, newColumn
}

func isTree(row int, column int, lines []string) bool {
	return lines[row][column] == '#'
}

func findTreesInSlope(down int, right int, lines []string) int {
	row, column := 0, 0 // start position
	trees := 0
	rows := len(lines)
	columns := len(lines[0])

	for row < rows-1 {
		row, column = getNextPosition(row, column, columns, down, right)
		if isTree(row, column, lines) {
			trees++
		}
	}

	return trees
}

func day3_1(lines []string) int {
	return findTreesInSlope(1, 3, lines)
}

func day3_2(lines []string) int {
	/*
		Right 1, down 1.
		Right 3, down 1. (This is the slope you already checked.)
		Right 5, down 1.
		Right 7, down 1.
		Right 1, down 2.
		And multiply
	*/

	slopes := []struct {
		right int
		down  int
	}{
		{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2},
	}
	result := 1

	for _, slope := range slopes {
		trees := findTreesInSlope(slope.down, slope.right, lines)
		result *= trees
	}

	return result
}

func day3Main() {
	lines, err := ReadLines("../inputs/day3")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	count := day3_1(lines)
	fmt.Println(count)

	count2 := day3_2(lines)
	fmt.Println(count2)
}
