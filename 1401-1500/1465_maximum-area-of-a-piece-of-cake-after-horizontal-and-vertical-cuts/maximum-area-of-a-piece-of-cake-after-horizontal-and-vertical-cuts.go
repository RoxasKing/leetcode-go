package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Greedy
// Sorting

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	m, n := len(horizontalCuts), len(verticalCuts)
	r := max(verticalCuts[0], w-verticalCuts[n-1])
	c := max(horizontalCuts[0], h-horizontalCuts[m-1])
	for i := 0; i < n-1; i++ {
		r = max(r, verticalCuts[i+1]-verticalCuts[i])
	}
	for i := 0; i < m-1; i++ {
		c = max(c, horizontalCuts[i+1]-horizontalCuts[i])
	}
	return r * c % (1e9 + 7)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
