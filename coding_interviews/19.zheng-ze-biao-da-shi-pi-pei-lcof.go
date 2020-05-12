package codinginterviews

/*
  请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for j := 1; j < len(p) && dp[0][j-1]; j++ {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}
	for i := range s {
		for j := range p {
			if s[i] == p[j] || '.' == p[j] {
				dp[i+1][j+1] = dp[i][j]
			} else if '*' == p[j] {
				dp[i+1][j+1] = dp[i+1][j-1]
				if !dp[i+1][j+1] && (s[i] == p[j-1] || '.' == p[j-1]) {
					dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
