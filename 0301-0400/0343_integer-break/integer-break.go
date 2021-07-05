package main

// Tags:
// Dynamic Programming
func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = Max(dp[i], j*Max(i-j, dp[i-j]))
		}
	}
	return dp[n]
}

// Dynamic Programming
func integerBreak2(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = Max(2*Max(i-2, dp[i-2]), 3*Max(i-3, dp[i-3]))
	}
	return dp[n]
}

func Max(a, b int) int {
	if a < b {
		return a
	}
	return b
}
