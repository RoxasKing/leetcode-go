package main

/*
  Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

  Example 1:
    Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
    Output: 4

  Example 2:
    Input: matrix = [["0","1"],["1","0"]]
    Output: 1

  Example 3:
    Input: matrix = [["0"]]
    Output: 0

  Constraints:
    1. m == matrix.length
    2. n == matrix[i].length
    3. 1 <= m, n <= 300
    4. matrix[i][j] is '0' or '1'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximal-square
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func maximalSquare(matrix [][]byte) int {
	side := 0
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[1][j] = int(matrix[i][j] - '0')
			if dp[1][j] == 0 {
				continue
			}
			if i > 0 && j > 0 {
				dp[1][j] += Min(dp[0][j-1], Min(dp[0][j], dp[1][j-1]))
			}
			side = Max(side, dp[1][j])
		}
		dp[0], dp[1] = dp[1], dp[0]
	}
	return side * side
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
