package codinginterviews

/*
  我们把只包含因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

  说明:
    1 是丑数。
    n 不超过1690。
*/

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	mul2, mul3, mul5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = Min(Min(dp[mul2]*2, dp[mul3]*3), dp[mul5]*5)
		if dp[mul2]*2 <= dp[i] {
			mul2++
		}
		if dp[mul3]*3 <= dp[i] {
			mul3++
		}
		if dp[mul5]*5 <= dp[i] {
			mul5++
		}
	}
	return dp[n-1]
}
