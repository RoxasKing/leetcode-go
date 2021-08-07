package main

import "sort"

// Tags:
// Two Pointers

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	out := 0
	for i := len(nums) - 1; i >= 2; i-- {
		l, r := 0, i-1
		for l < r {
			if nums[l]+nums[r] > nums[i] {
				out += r - l
				r--
			} else {
				l++
			}
		}
	}
	return out
}
