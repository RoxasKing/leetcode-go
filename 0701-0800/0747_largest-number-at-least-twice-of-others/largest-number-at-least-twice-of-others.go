package main

// Difficulty:
// Easy

func dominantIndex(nums []int) int {
	a, b := -1, -1
	out := -1
	for i, num := range nums {
		if num > b {
			a, b = b, num
			out = i
		} else if num > a {
			a = num
		}
	}
	if a == 0 || a == -1 || b/a >= 2 {
		return out
	}
	return -1
}
