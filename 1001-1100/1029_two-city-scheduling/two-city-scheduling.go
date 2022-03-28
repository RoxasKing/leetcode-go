package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Greedy
// Sorting

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool { return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1] })
	n := len(costs) >> 1
	out := 0
	for i := 0; i < n; i++ {
		out += costs[i][0] + costs[i+n][1]
	}
	return out
}
