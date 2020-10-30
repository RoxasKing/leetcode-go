package main

/*
  给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

  你可以对一个单词进行如下三种操作：
    插入一个字符
    删除一个字符
    替换一个字符

  示例 1：
    输入：word1 = "horse", word2 = "ros"
    输出：3
    解释：
    horse -> rorse (将 'h' 替换为 'r')
    rorse -> rose (删除 'r')
    rose -> ros (删除 'e')

  示例 2：
    输入：word1 = "intention", word2 = "execution"
    输出：5
    解释：
    intention -> inention (删除 't')
    inention -> enention (将 'i' 替换为 'e')
    enention -> exention (将 'n' 替换为 'x')
    exention -> exection (将 'n' 替换为 'c')
    exection -> execution (插入 'u')

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/edit-distance
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := range word1 {
		for j := range word2 {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j] // no-op
			} else {
				dp[i+1][j+1] = Min(dp[i][j], Min(dp[i+1][j], dp[i][j+1])) + 1 // exchange/add/del
			}
		}
	}
	return dp[m][n]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
