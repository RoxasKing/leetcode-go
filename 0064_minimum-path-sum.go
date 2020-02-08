package leetcode

/*
  给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
  说明：每次只能向下或者向右移动一步。
*/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	dp := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if j > 0 && (i == 0 || dp[j-1] < dp[j]) {
				dp[j] = dp[j-1] + grid[i][j]
			} else {
				dp[j] += grid[i][j]
			}
		}
	}
	return dp[len(dp)-1]
}

func minPathSum2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for i := range grid {
		for j := range grid[i] {
			if i == 0 && j > 0 ||
				i > 0 && j > 0 && grid[i][j-1] < grid[i-1][j] {
				grid[i][j] += grid[i][j-1]
			} else if i > 0 {
				grid[i][j] += grid[i-1][j]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
