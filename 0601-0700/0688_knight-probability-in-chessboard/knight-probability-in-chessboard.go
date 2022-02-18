package main

// Difficulty:
// Medium

// Tags:
// DFS
// Memoization
// Dynamic Programming

func knightProbability(n int, k int, row int, column int) float64 {
	dr := []int{-2, -1, 1, 2, 2, 1, -1, -2}
	dc := []int{1, 2, 2, 1, -1, -2, -2, -1}
	dp := make([][][]float64, n)
	for i := range dp {
		dp[i] = make([][]float64, n)
		for j := range dp[i] {
			dp[i][j] = make([]float64, k+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1.0
			}
		}
	}
	var dfs func(int, int, int) float64
	dfs = func(x, y, k int) float64 {
		if x < 0 || n-1 < x || y < 0 || n-1 < y {
			return 0.0
		}
		if k == 0 {
			return 1.0
		}
		if dp[x][y][k] != -1.0 {
			return dp[x][y][k]
		}
		out := 0.0
		for i := 0; i < 8; i++ {
			r, c := x+dr[i], y+dc[i]
			out += dfs(r, c, k-1)
		}
		dp[x][y][k] = out / 8.0
		return dp[x][y][k]
	}
	return dfs(row, column, k)
}
