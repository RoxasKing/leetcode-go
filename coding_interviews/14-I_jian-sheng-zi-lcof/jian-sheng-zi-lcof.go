package main

// Tags:
// Dynamic Programming
func cuttingRope(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= n; i++ {
		for j := 1; j <= i>>1; j++ {
			dp[i] = Max(dp[i], dp[i-j]*j)
		}
	}
	return dp[n]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
