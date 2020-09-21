package main

/*
  给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

  说明：
    分隔时可以重复使用字典中的单词。
    你可以假设字典中没有重复的单词。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/word-break-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func wordBreak(s string, wordDict []string) []string {
	if !isValid(s, wordDict) {
		return nil
	}
	n := len(s)
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			if i < len(word) || s[i-len(word):i] != word || len(dp[i-len(word)]) == 0 {
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

func isValid(s string, wordDict []string) bool {
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
