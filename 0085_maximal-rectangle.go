package main

/*
  给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
*/

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var maxArea int
	dp := make([]int, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == '1' {
				dp[j] = dp[j] + 1
			} else {
				dp[j] = 0
			}
		}
		maxArea = Max(maxArea, largestRectangleArea(dp))
	}
	return maxArea
}
