package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting
// Hash
// Two Pointers

func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	n := len(nums1)
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool { return nums2[ids[i]] < nums2[ids[j]] })
	left := []int{}
	o := make([]int, n)
	k := 0
	for i := 0; i < n && k < n; i, k = i+1, k+1 {
		for ; i < n && nums1[i] <= nums2[ids[k]]; i++ {
			left = append(left, nums1[i])
		}
		if i == n {
			break
		}
		o[ids[k]] = nums1[i]
	}
	for i := 0; k < n; i, k = i+1, k+1 {
		o[ids[k]] = left[i]
	}
	return o
}
