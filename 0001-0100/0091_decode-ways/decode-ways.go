package main

// Tags:
// Dynamic Programming
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := range s {
		if s[i] != '0' {
			dp[i+1] += dp[i]
		}
		if i > 0 && (s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6') {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[n]
}
