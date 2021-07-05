package main

// Tags:
// Dynamic Programming
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := range word1 {
		for j := range word2 {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j] // no-op
			} else {
				dp[i+1][j+1] = Min(dp[i][j], Min(dp[i+1][j], dp[i][j+1])) + 1 // exchange/add/del
			}
		}
	}
	return dp[m][n]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
