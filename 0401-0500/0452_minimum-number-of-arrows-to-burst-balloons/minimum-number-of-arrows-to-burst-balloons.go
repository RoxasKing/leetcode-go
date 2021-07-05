package main

import "sort"

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})
	r := points[0][1]
	count := 1
	for i := 1; i < n; i++ {
		if points[i][0] > r {
			r = points[i][1]
			count++
		} else if points[i][1] < r {
			r = points[i][1]
		}
	}
	return count
}
