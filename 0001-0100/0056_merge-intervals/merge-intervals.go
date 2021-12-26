package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	out := [][]int{}
	tail := -1
	for _, e := range intervals {
		if tail < 0 || out[tail][1] < e[0] {
			out = append(out, e)
			tail++
		} else if out[tail][1] < e[1] {
			out[tail][1] = e[1]
		}
	}
	return out
}
