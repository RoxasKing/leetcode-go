package leetcode

/*
  给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
  说明：
    拆分时可以重复使用字典中的单词。
    你可以假设字典中没有重复的单词。
*/

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}
	dp := make([]int, len(s))
	for i := range dp {
		dp[i] = len(s)
	}
	for i := range s {
		dp[i] = len(s)
		if dict[s[0:i+1]] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if dict[s[j+1:i+1]] && dp[i] > dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[len(s)-1] < len(s)
}

func wordBreak2(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	for i := range s {
		for _, w := range wordDict {
			if i < len(w)-1 {
				continue
			}
			if string(s[i-len(w)+1:i+1]) == w &&
				(i == len(w)-1 || dp[i-len(w)]) {
				dp[i] = true
			}
		}
	}
	return dp[len(s)-1]
}
