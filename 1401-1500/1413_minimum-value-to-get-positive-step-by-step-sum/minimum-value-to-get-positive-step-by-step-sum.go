package main

// Difficulty:
// Easy

func minStartValue(nums []int) int {
	cur, out := 0, 1<<31-1
	for _, num := range nums {
		cur += num
		if cur < out {
			out = cur
		}
	}
	if out = 1 - out; out <= 0 {
		out = 1
	}
	return out
}
