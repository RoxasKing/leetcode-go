package main

// Tags:
// Math + Dynamic Programming
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = Min(dp[i2]*2, Min(dp[i3]*3, dp[i5]*5))
		if dp[i2]*2 == dp[i] {
			i2++
		}
		if dp[i3]*3 == dp[i] {
			i3++
		}
		if dp[i5]*5 == dp[i] {
			i5++
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
