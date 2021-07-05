package main

// Tags:
// Dynamic Programming
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < n; i++ {
		dp[i] = cost[i] + Min(dp[i-1], dp[i-2])
	}
	return Min(dp[n-2], dp[n-1])
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
