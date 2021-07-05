package main

// Tags:
// Dynamic Programming
func divisorGame(n int) bool {
	dp := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// Math
func divisorGame2(N int) bool {
	return N&1 == 0
}
