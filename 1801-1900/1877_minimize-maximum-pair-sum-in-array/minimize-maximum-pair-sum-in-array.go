package main

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	out := 0
	for i := 0; i < n>>1; i++ {
		out = Max(out, nums[i]+nums[n-1-i])
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
