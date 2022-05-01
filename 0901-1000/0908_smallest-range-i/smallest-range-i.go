package main

// Difficulty:
// Easy

func smallestRangeI(nums []int, k int) int {
	max, min := -1<<31, 1<<31-1
	for _, x := range nums {
		if max < x {
			max = x
		}
		if min > x {
			min = x
		}
	}
	if max-min <= k*2 {
		return 0
	}
	return max - min - k*2
}
