package leetcode

/*
  给定一个字符串 S 和一个字符串 T，计算在 S 的子序列中 T 出现的个数。
  一个字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
*/

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	for i := range dp {
		dp[i][0] = 1
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			dp[i][j] = dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	return dp[len(s)][len(t)]
}

func numDistinct2(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}
	dp := make([]int, len(t))
	for i := range s {
		for j := len(t) - 1; j >= 0; j-- {
			if s[i] == t[j] {
				if j > 0 {
					dp[j] += dp[j-1]
				} else {
					dp[j]++
				}
			}
		}
	}
	return dp[len(t)-1]
}
