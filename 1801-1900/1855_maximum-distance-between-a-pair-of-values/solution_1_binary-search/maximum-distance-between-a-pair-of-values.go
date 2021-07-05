package main

import (
	"sort"
)

// Tags:
// Binary Search
func maxDistance(nums1 []int, nums2 []int) int {
	out := 0
	m, n := len(nums1), len(nums2)
	for i := 0; i < m; i++ {
		j := sort.Search(n, func(j int) bool { return nums2[j] < nums1[i] }) - 1
		if j < i {
			continue
		}
		out = Max(out, j-i)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
