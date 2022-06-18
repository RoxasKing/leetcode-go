package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Sorting

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	o := 0
	for i := range expected {
		if expected[i] != heights[i] {
			o++
		}
	}
	return o
}
