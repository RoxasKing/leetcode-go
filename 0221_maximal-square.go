package leetcode

/*
  在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
*/

func maximalSquare(matrix [][]byte) int {
	var side int
	dp := make([][]int, len(matrix))
	for i := range matrix {
		dp[i] = make([]int, len(matrix[0]))
		for j := range matrix[0] {
			if matrix[i][j] != '1' {
				continue
			}
			dp[i][j] = int(matrix[i][j] - '0')
			if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] += Min(Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])
			}
			side = Max(side, dp[i][j])
		}
	}
	return side * side
}
