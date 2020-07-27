package leetcode

/*
  给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
  你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
  字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

  后续挑战 :
    如果有大量输入的 S，称作S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/is-subsequence
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Double Pointer
func isSubsequence(s string, t string) bool {
	for s != "" && t != "" {
		if s[0] == t[0] {
			s = s[1:]
		}
		t = t[1:]
	}
	return s == ""
}

// Dynamic Programming
func isSubsequence2(s string, t string) bool {
	dp := make([][26]int, len(t)+1)
	for i := 0; i < 26; i++ {
		dp[len(t)][i] = len(t)
	}
	for i := len(t) - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j+'a') {
				dp[i][j] = i
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	var index int
	for i := 0; i < len(s); i++ {
		if dp[index][int(s[i]-'a')] == len(t) {
			return false
		}
		index = dp[index][int(s[i]-'a')] + 1
	}
	return true
}
