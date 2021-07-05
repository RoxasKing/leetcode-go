package main

import "sort"

// Sort

func reductionOperations(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	lastNum := nums[len(nums)-1]
	out := 0
	for i := 1; nums[i-1] != lastNum; i++ {
		if nums[i] < nums[i-1] {
			out += i
		}
	}
	return out
}
