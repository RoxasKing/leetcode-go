package main

import (
	"sort"
)

// Tags:
// Binary Search
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	idxs := make([]int, n)
	for i := 0; i < n; i++ {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return intervals[idxs[i]][0] < intervals[idxs[j]][0]
	})

	out := make([]int, n)
	for i, interval := range intervals {
		j := sort.Search(n, func(j int) bool {
			return intervals[idxs[j]][0] >= interval[1]
		})

		if j == n {
			out[i] = -1
		} else {
			out[i] = idxs[j]
		}
	}
	return out
}
