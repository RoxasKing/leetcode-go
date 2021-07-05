package main

import "sort"

// Dynamic Programming
func canCross(stones []int) bool {
	dict := make(map[int]int)
	for i, stone := range stones {
		dict[stone] = i
	}
	n := len(stones)
	dp := make([][]int, n)
	dp[0] = append(dp[0], 0)
	for i := 0; i < n-1; i++ {
		sort.Ints(dp[i])
		for j, k := range dp[i] {
			if j > 0 && k == dp[i][j-1] {
				continue
			}
			if next, ok := dict[stones[i]+k-1]; ok {
				dp[next] = append(dp[next], k-1)
			}
			if next, ok := dict[stones[i]+k]; ok {
				dp[next] = append(dp[next], k)
			}
			if next, ok := dict[stones[i]+k+1]; ok {
				dp[next] = append(dp[next], k+1)
			}
		}
	}
	return len(dp[n-1]) > 0
}
