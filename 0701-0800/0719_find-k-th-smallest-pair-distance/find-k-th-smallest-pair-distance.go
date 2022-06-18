package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Binary Search
// Two Pointers

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r := 0, int(1e6)
	for l < r {
		m := (l + r) / 2
		c := 0
		for i, j := 0, 1; j < n; j++ {
			for ; i < n && nums[j]-nums[i] > m; i++ {
			}
			c += j - i
		}
		if c < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
