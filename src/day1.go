package main

import (
	"fmt"
	"log"
)

func createSet(values []int) map[int]struct{} {
	set := make(map[int]struct{})
	for _, n := range values {
		set[n] = struct{}{}
	}
	return set
}

func day1_1(values []int, target int) (int, int) {
	// Create a set
	set := createSet(values)
	fmt.Println(set)
	for n := range set {
		var d = target - n
		if _, ok := set[d]; ok {
			return n, d
		}
	}
	return 0, 0
}

func day1_2(values []int, target int) (int, int, int) {
	set := createSet(values)

	for i, n1 := range values {
		for _, n2 := range values[i:] {
			d := target - (n1 + n2)
			if _, ok := set[d]; ok {
				return n1, n2, d
			}
		}
	}

	return 0, 0, 0
}

func day1Main() {
	lines, err := ReadLines("../inputs/day1")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	values := ConvertLinesToInt(lines)
	// fmt.Println(values)
	n1, n2 := day1_1(values, 2020)
	fmt.Println(n1 * n2)

	n1, n2, n3 := day1_2(values, 2020)
	fmt.Println(n1, n2, n3)
	fmt.Println(n1 * n2 * n3)
	//11581560
}
