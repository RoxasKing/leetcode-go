package main

import (
	"sort"
)

// Tags:
// Dynamic Programming
func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	m := len(cuts) + 2
	lens := make([]int, m)
	copy(lens[1:m-1], cuts)
	lens[m-1] = n
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for k := 2; k < m; k++ {
		for i := 0; i < m-k; i++ {
			dp[i][i+k] = 1<<31 - 1
			for j := 1; j < k; j++ {
				dp[i][i+k] = Min(dp[i][i+k], lens[i+k]-lens[i]+dp[i][i+j]+dp[i+j][i+k])
			}
		}
	}
	return dp[0][m-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
