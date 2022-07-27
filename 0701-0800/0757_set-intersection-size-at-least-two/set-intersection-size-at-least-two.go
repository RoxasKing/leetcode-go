package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Sorting
// Greedy

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	o, one, two := 0, -1, -1
	for _, e := range intervals {
		if e[0] <= one {
			continue
		}
		if two < e[0] {
			one, two = e[1]-1, e[1]
			o += 2
			continue
		}
		if two == e[1] {
			one = e[1] - 1
		} else {
			one, two = two, e[1]
		}
		o++
	}
	return o
}
