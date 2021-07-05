package main

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	out := 0
	for i := 0; i < len(nums); i += 2 {
		out += Min(nums[i], nums[i+1])
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
