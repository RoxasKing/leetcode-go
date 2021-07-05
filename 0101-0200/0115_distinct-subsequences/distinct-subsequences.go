package main

// Tags:
// Dynamic Programming
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = 1
	}
	for i := range t {
		for j := range s {
			dp[i+1][j+1] = dp[i+1][j]
			if t[i] == s[j] {
				dp[i+1][j+1] += dp[i][j]
			}
		}
	}
	return dp[n][m]
}
