package main

// Difficulty:
// Easy

func maximumDifference(nums []int) int {
	out, min := -1, 1<<31-1
	for _, num := range nums {
		if num > min && out < num-min {
			out = num - min
		}
		if min > num {
			min = num
		}
	}
	return out
}
