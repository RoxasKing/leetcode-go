package main

// Tags:
// Dynamic Programming
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	copy(dp, triangle[n-1])
	for i := n - 2; i >= 0; i-- {
		for j := range triangle[i] {
			dp[j] = Min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
