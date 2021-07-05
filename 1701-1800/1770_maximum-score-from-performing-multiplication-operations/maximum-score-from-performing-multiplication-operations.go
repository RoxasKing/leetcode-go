package main

// Tags:
// Dynamic Programming
func maximumScore(nums []int, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + nums[i-1]*multipliers[i-1]
		dp[0][i] = dp[0][i-1] + nums[n-i]*multipliers[i-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; i+j <= m; j++ {
			k := i + j - 1
			dp[i][j] = Max(dp[i][j-1]+nums[n-j]*multipliers[k], dp[i-1][j]+nums[i-1]*multipliers[k])
		}
	}

	out := -1 << 31
	for i := 0; i <= m; i++ {
		out = Max(out, dp[i][m-i])
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
