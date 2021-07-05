package main

// Tags:
// Dynamic Programming
func superEggDrop(k int, n int) int {
	dp := make([][]int, n+1) // j个 鸡蛋操作 i次 可以确定的层数
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i < n; i++ { // 操作次数
		for j := 1; j <= k; j++ { // 鸡蛋数
			dp[i][j] = dp[i-1][j-1] + 1 + dp[i-1][j]
			if dp[i][j] >= n {
				return i
			}
		}
	}
	return n
}
