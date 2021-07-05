package main

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	index := sort.Search(n, func(i int) bool { return intervals[i][0] >= newInterval[0] })
	out := make([][]int, 0, n+1)
	if index == 0 {
		out = append(out, newInterval)
	} else {
		out = append(out, intervals[:index]...)
		index--
		intervals[index] = newInterval
	}
	for i := index; i < n; i++ {
		if intervals[i][0] > out[len(out)-1][1] {
			out = append(out, intervals[i])
		} else if intervals[i][1] > out[len(out)-1][1] {
			out[len(out)-1][1] = intervals[i][1]
		}
	}
	return out
}
