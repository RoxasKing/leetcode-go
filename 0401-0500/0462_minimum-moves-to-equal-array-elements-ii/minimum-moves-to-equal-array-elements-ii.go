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
	psum := make([]int, n)
	psum[0] = nums[0]
	for i := 1; i < n; i++ {
		psum[i] = psum[i-1] + nums[i]
	}
	o := 1<<31 - 1
	for i := 0; i < n; i++ {
		d := 0
		if i > 0 {
			d += i*nums[i] - psum[i-1]
		}
		if i+1 < n {
			d += psum[n-1] - psum[i] - (n-1-i)*nums[i]
		}
		if o > d {
			o = d
		}
	}
	return o
}
