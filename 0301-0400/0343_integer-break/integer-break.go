package main

/*
  给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

  示例 1:
    输入: 2
    输出: 1
    解释: 2 = 1 + 1, 1 × 1 = 1。

  示例 2:
    输入: 10
    输出: 36
    解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/integer-break
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = Max(dp[i], j*Max(i-j, dp[i-j]))
		}
	}
	return dp[n]
}

// Dynamic Programming
func integerBreak2(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = Max(2*Max(i-2, dp[i-2]), 3*Max(i-3, dp[i-3]))
	}
	return dp[n]
}

func Max(a, b int) int {
	if a < b {
		return a
	}
	return b
}
