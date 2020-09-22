package main

/*
  给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
  说明：
    拆分时可以重复使用字典中的单词。
    你可以假设字典中没有重复的单词。
*/

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
