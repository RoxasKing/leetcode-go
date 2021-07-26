package main

import "sort"

// Greedy
// Prefix Sum
// Binary Search

func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)
	pSum := make([]int, n)
	pSum[0] = chalk[0]
	for i := 1; i < n; i++ {
		pSum[i] = pSum[i-1] + chalk[i]
	}
	k %= pSum[n-1]
	return sort.Search(n, func(i int) bool { return pSum[i] > k })
}
