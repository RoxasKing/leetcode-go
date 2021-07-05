package main

// Tags:
// Dynamic Programming
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp := 1<<31 - 1
			if i > 0 && dp[j] < tmp {
				tmp = dp[j]
			}
			if j > 0 && dp[j-1] < tmp {
				tmp = dp[j-1]
			}
			if tmp == 1<<31-1 {
				tmp = 0
			}
			dp[j] = tmp + grid[i][j]
		}
	}
	return dp[n-1]
}
