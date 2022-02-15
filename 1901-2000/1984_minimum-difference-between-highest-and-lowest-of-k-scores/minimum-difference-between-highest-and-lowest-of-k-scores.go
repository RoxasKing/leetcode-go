package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Sorting

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	out := 1<<31 - 1
	for i := 0; i+k-1 < len(nums); i++ {
		if diff := nums[i+k-1] - nums[i]; diff < out {
			out = diff
		}
	}
	return out
}
