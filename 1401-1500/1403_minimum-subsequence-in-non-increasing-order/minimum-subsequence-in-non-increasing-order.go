package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Greedy
// Sorting

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	sum := 0
	for _, x := range nums {
		sum += x
	}
	i := 0
	for x := 0; i < len(nums) && x <= sum-x; i++ {
		x += nums[i]
	}
	return nums[:i]
}
