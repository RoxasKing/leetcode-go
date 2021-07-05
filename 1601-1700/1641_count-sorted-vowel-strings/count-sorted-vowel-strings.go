package main

// Tags:
// Dynamic Programming
func countVowelStrings(n int) int {
	dp := [2][5]int{}
	for i := 0; i < 5; i++ {
		dp[0][i] = i + 1
	}
	for i := 1; i < n; i++ {
		dp[1][0] = 1
		for j := 1; j < 5; j++ {
			dp[1][j] = dp[1][j-1] + dp[0][j]
		}
		dp[0], dp[1] = dp[1], dp[0]
	}
	return dp[0][4]
}
