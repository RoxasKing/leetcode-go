package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting
// Prefix Sum

func minMoves2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	pSum := make([]int, n+1)
	for i, x := range nums {
		pSum[i+1] = pSum[i] + x
	}
	o := pSum[n] - nums[0]*n
	for i, x := range nums {
		if t := (x*i - pSum[i]) + (pSum[n] - pSum[i] - (n-i)*x); o > t {
			o = t
		}
	}
	return o
}
