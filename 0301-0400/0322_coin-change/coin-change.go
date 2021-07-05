package main

import "sort"

// Dynamic Programming
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = 1<<31 - 1
		for _, coin := range coins {
			if coin > i {
				break
			}
			dp[i] = Min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == 1<<31-1 {
		return -1
	}
	return dp[amount]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
