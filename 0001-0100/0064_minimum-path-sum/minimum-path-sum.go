package main

/*
  给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

  说明：每次只能向下或者向右移动一步。

  示例:
    输入:
    [
      [1,3,1],
      [1,5,1],
      [4,2,1]
    ]
    输出: 7
    解释: 因为路径 1→3→1→1→1 的总和最小。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-path-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] += grid[i][j]
			} else {
				dp[j] = Min(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
