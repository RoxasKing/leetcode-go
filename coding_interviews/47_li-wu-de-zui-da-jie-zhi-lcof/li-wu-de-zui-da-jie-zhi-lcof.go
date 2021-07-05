package main

// Tags:
// Dynamic Programming
func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = Max(dp[j], dp[j-1]) + grid[i][j]
		}
	}
	return dp[n-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
