package main

// Tags:
// Dynamic Programming
func longestPalindrome(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	k := m + n
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	out := 0
	word := word1 + word2
	for i := 0; i < k; i++ {
		for j := 0; i+j < k; j++ {
			if word[i] == word[k-1-j] {
				if dp[0][j] != 0 {
					dp[1][j+1] = dp[0][j] + 1
					if i < k-1-j {
						dp[1][j+1]++
					}
				} else if i < m && j < n { // 关键,左端在word1之中，右端在word2之中
					dp[1][j+1] = 2
				}
				out = Max(out, dp[1][j+1])
			} else {
				dp[1][j+1] = Max(dp[1][j], dp[0][j+1])
			}
		}
		dp[0], dp[1] = dp[1], dp[0]
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
