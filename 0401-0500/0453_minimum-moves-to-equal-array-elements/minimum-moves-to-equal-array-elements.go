package main

// Difficulty:
// Easy

// Tags:
// Math

func minMoves(nums []int) int {
	min := int(1e9)
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	out := 0
	for _, num := range nums {
		out += num - min
	}
	return out
}
