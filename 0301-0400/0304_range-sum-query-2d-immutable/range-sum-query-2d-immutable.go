package main

// Tags:
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
