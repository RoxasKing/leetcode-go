package main

import "sort"

// DifficultY:
// Medium

// Tags:
// Sorting

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return a[1] > b[1]
	})
	out := 0
	for k, i := -1, 0; i < len(intervals); i++ {
		if k == -1 || intervals[k][1] < intervals[i][1] {
			k = i
			out++
		}
	}
	return out
}
