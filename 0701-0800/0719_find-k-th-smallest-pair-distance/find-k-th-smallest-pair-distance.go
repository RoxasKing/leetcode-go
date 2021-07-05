package main

import "sort"

// Binary Search
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r := 0, int(1e6)
	for l < r {
		dist := (l + r) >> 1
		if countSmaller(nums, n, dist) < k {
			l = dist + 1
		} else {
			r = dist
		}
	}
	return l
}

func countSmaller(nums []int, n, dist int) int {
	out := 0
	for l, r := 0, 1; r < n; r++ {
		for nums[r]-nums[l] > dist {
			l++
		}
		out += r - l
	}
	return out
}
