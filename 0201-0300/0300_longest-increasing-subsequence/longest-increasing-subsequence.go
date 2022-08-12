package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Binary Search

func lengthOfLIS(nums []int) int {
	a := []int{}
	for _, x := range nums {
		if i := sort.Search(len(a), func(i int) bool { return a[i] >= x }); i < len(a) {
			a[i] = x
			continue
		}
		a = append(a, x)
	}
	return len(a)
}
