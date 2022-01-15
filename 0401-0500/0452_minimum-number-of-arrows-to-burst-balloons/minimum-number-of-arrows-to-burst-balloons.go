package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting
// Greedy

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })
	out := 1
	for i, r := 1, points[0][1]; i < n; i++ {
		if points[i][0] > r {
			r = points[i][1]
			out++
		} else if points[i][1] < r {
			r = points[i][1]
		}
	}
	return out
}
