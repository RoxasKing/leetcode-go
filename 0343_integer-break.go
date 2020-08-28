package main

/*
  给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

  说明: 你可以假设 n 不小于 2 且不大于 58。
*/

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = Max(dp[i], j*Max(i-j, dp[i-j]))
		}
	}
	return dp[n]
}

func integerBreak2(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = Max(2*Max(i-2, dp[i-2]), 3*Max(i-3, dp[i-3]))
	}
	return dp[n]
}
