package main

// Tags:
// Dynamic Programming
func wordBreak(s string, wordDict []string) []string {
	n := len(s)
	if !canSplit(s, n, wordDict) {
		return nil
	}
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			if i < len(word) || s[i-len(word):i] != word {
				continue
			}
			for _, str := range dp[i-len(word)] {
				if str != "" {
					str += " "
				}
				dp[i] = append(dp[i], str+word)
			}
		}
	}
	return dp[n]
}

func canSplit(s string, n int, wordDict []string) bool {
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
