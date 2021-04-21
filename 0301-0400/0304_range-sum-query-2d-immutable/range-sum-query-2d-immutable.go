package main

/*
  Given a 2D matrix matrix, find the sum of the elements inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).

  The above rectangle (with the red border) is defined by (row1, col1) = (2, 1) and (row2, col2) = (4, 3), which contains sum = 8.

  Example:
    Given matrix = [
      [3, 0, 1, 4, 2],
      [5, 6, 3, 2, 1],
      [1, 2, 0, 1, 5],
      [4, 1, 0, 1, 7],
      [1, 0, 3, 0, 5]
    ]
    sumRegion(2, 1, 4, 3) -> 8
    sumRegion(1, 1, 2, 2) -> 11
    sumRegion(1, 2, 2, 4) -> 12

  Note:
    1. You may assume that the matrix does not change.
    2. There are many calls to sumRegion function.
    3. You may assume that row1 ≤ row2 and col1 ≤ col2.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/range-sum-query-2d-immutable
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = matrix[i][j]
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
			if i > 0 && j > 0 {
				dp[i][j] -= dp[i-1][j-1]
			}
		}
	}
	return NumMatrix{dp: dp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	dp := this.dp
	out := dp[row2][col2]
	if row1 > 0 {
		out -= dp[row1-1][col2]
	}
	if col1 > 0 {
		out -= dp[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		out += dp[row1-1][col1-1]
	}
	return out
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
