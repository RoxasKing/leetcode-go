package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Sort

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	min := 1<<31 - 1
	out := 0
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
		if nums[i] < min {
			min = nums[i]
		}
		out += nums[i]
	}
	if k&1 == 1 {
		out -= min << 1
	}
	return out
}
