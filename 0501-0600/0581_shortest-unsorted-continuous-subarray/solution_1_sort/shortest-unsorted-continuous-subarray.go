package main

import "sort"

// Tags:
// Sort

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)
	l, r := 0, n-1
	for l < n && sorted[l] == nums[l] {
		l++
	}
	for r >= l && sorted[r] == nums[r] {
		r--
	}
	return r + 1 - l
}
