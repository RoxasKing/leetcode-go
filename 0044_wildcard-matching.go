package main

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
  '?' 可以匹配任何单个字符。
  '*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。
说明:
  s 可能为空，且只包含从 a-z 的小写字母。
  p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
*/

// Dynamic Programming
func isMatchII(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(p) && p[i-1] == '*'; i++ {
		dp[0][i] = true
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

// Greedy algorithm
func isMatchII2(s string, p string) bool {
	for len(s) > 0 && len(p) > 0 && p[len(p)-1] != '*' {
		if s[len(s)-1] == p[len(p)-1] || '?' == p[len(p)-1] {
			s = s[:len(s)-1]
			p = p[:len(p)-1]
		} else {
			return false
		}
	}
	if s != "" && p == "" {
		return false
	}
	sIndex, pIndex := 0, 0
	sRecord, pRecord := -1, -1
	for sIndex < len(s) && pIndex < len(p) {
		if s[sIndex] == p[pIndex] || '?' == p[pIndex] {
			sIndex++
			pIndex++
		} else if '*' == p[pIndex] {
			pIndex++
			sRecord, pRecord = sIndex, pIndex
		} else if sRecord != -1 && sRecord+1 < len(s) {
			sRecord++
			sIndex, pIndex = sRecord, pRecord
		} else {
			return false
		}
	}
	for pIndex < len(p) {
		if p[pIndex] != '*' {
			return false
		}
		pIndex++
	}
	return true
}
