package main

/*
  Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

  Note: You can only move either down or right at any point in time.

  Example 1:
    Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
    Output: 7
    Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

  Example 2:
    Input: grid = [[1,2,3],[4,5,6]]
    Output: 12

  Constraints:
    1. m == grid.length
    2. n == grid[i].length
    3. 1 <= m, n <= 200
    4. 0 <= grid[i][j] <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-path-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
