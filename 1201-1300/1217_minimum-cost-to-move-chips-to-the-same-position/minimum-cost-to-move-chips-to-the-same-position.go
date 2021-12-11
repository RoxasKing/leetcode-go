package main

// Difficulty:
// Easy

// Tags:
// Math
// Greedy

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, pos := range position {
		if pos&1 == 1 {
			odd++
		} else {
			even++
		}
	}
	if odd < even {
		return odd
	}
	return even
}
