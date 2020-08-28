package main

/*
  给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。
  你可以对一个单词进行如下三种操作：
    插入一个字符
    删除一个字符
    替换一个字符
*/

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i-1][j-1], Min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
