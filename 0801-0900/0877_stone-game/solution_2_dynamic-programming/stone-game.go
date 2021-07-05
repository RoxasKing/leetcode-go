package main

// Tags:
// Dynamic Programming

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([]int, n)
	copy(dp, piles)
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i] = Max(piles[i]-dp[j], piles[j]-dp[j-1])
		}
	}
	return dp[n-1] >= 0
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
