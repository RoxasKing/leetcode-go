package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting
// Binary Search

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	h := make([]int, n)
	for i := range h {
		h[i] = i
	}
	sort.Slice(h, func(i, j int) bool { return intervals[h[i]][0] < intervals[h[j]][0] })
	o := make([]int, n)
	for i, e := range intervals {
		o[i] = -1
		if j := sort.Search(n, func(j int) bool { return intervals[h[j]][0] >= e[1] }); j < n {
			o[i] = h[j]
		}
	}
	return o
}
