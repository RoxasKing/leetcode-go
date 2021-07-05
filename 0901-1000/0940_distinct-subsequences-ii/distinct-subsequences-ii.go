package main

// Tags:
// Dynamic Programming
func distinctSubseqII(S string) int {
	mod := int(1e9 + 7)
	n := len(S)
	dp := make([]int, n+1)
	dp[0] = 1
	last := [26]int{}
	for i := range last {
		last[i] = -1
	}
	for i := 0; i < n; i++ {
		letter := S[i] - 'a'
		dp[i+1] = (dp[i] << 1) % mod
		if last[letter] != -1 {
			dp[i+1] = (dp[i+1] - dp[last[letter]] + mod) % mod
		}
		last[letter] = i
	}
	dp[n]--
	return dp[n]
}
