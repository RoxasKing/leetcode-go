package main

// Tags:
// Dynamic Programming
func maxCoins(nums []int) int {
	n := len(nums)
	vals := make([]int, n+2)
	copy(vals[1:n+1], nums)
	vals[0], vals[n+1] = 1, 1
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for l := n - 1; l >= 0; l-- {
		for r := l + 2; r < n+2; r++ {
			for m := l + 1; m < r; m++ {
				tmp := vals[l]*vals[m]*vals[r] + dp[l][m] + dp[m][r]
				dp[l][r] = Max(dp[l][r], tmp)
			}
		}
	}
	return dp[0][n+1]
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
