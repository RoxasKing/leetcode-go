package main

import "sort"

// Greedy
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	count := 0
	pre := -1 << 31
	for _, interval := range intervals {
		if interval[0] < pre {
			count++
		} else {
			pre = interval[1]
		}
	}
	return count
}
