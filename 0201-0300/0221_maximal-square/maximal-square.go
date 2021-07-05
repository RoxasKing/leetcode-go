package main

// Tags:
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
