package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting
// Two Pointers

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	out := 0
	for l, r := 0, len(nums)-1; l < r; {
		sum := nums[l] + nums[r]
		if sum < k {
			l++
		} else if sum > k {
			r--
		} else {
			out++
			l++
			r--
		}
	}
	return out
}
