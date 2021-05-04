package main

/*
  Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

  You have the following three operations permitted on a word:
    1. Insert a character
    2. Delete a character
    3. Replace a character

  Example 1:
    Input: word1 = "horse", word2 = "ros"
    Output: 3
    Explanation:
      horse -> rorse (replace 'h' with 'r')
      rorse -> rose (remove 'r')
      rose -> ros (remove 'e')

  Example 2:
    Input: word1 = "intention", word2 = "execution"
    Output: 5
    Explanation:
      intention -> inention (remove 't')
      inention -> enention (replace 'i' with 'e')
      enention -> exention (replace 'n' with 'x')
      exention -> exection (replace 'n' with 'c')
      exection -> execution (insert 'u')

  Constraints:
    1. 0 <= word1.length, word2.length <= 500
    2. word1 and word2 consist of lowercase English letters.

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
