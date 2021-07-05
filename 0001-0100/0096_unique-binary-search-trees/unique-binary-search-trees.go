package main

// Tags:
// Dynamic Programming
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ { // number of nodes
		for j := 1; j <= i; j++ { // root node's index
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
