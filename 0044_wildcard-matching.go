package leetcode

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
  '?' 可以匹配任何单个字符。
  '*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。
说明:
  s 可能为空，且只包含从 a-z 的小写字母。
  p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
*/

func isMatch0044(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := range s {
		for j := range p {
			if s[i] == p[j] || p[j] == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1]
			}
		}
	}
	return dp[len(s)][len(p)]
}

func isMatch00442(s string, p string) bool {
	i, j, iStar, jStar := 0, 0, -1, -1
	for i < len(s) {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i, j = i+1, j+1
		} else if j < len(p) && p[j] == '*' {
			iStar, jStar = i, j
			j++
		} else if iStar >= 0 {
			iStar++
			i, j = iStar, jStar+1
		} else {
			return false
		}
	}
	for j < len(p) && p[j] == '*' {
		j++
	}
	return j == len(p)
}
