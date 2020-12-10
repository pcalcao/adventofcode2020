package main

import (
	"bufio"
	"os"
	"strconv"
)

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ConvertLinesToInt - Convert a list of strings to ints
func ConvertLinesToInt(lines []string) []int {
	var values []int
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		values = append(values, n)
	}
	return values
}
