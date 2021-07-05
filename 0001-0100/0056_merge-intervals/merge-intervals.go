package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i := 0
	for j := 1; j < len(intervals); j++ {
		if intervals[j][0] > intervals[i][1] {
			i++
			intervals[i] = intervals[j]
		} else if intervals[j][1] > intervals[i][1] {
			intervals[i][1] = intervals[j][1]
		}
	}

	return intervals[:i+1]
}
