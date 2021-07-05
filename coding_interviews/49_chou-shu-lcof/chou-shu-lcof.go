package main

// Tags:
// Dynamic Programming
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

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
