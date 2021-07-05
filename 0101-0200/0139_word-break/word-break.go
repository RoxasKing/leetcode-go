package main

// Tags:
// Dynamic Programming
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			if i >= len(word) && s[i-len(word):i] == word && dp[i-len(word)] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
