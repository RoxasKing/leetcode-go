package main

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
  '?' 可以匹配任何单个字符。
  '*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:
  s 可能为空，且只包含从 a-z 的小写字母。
  p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。

示例 1:
  输入:
  s = "aa"
  p = "a"
  输出: false
  解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:
  输入:
  s = "aa"
  p = "*"
  输出: true
  解释: '*' 可以匹配任意字符串。

示例 3:
  输入:
  s = "cb"
  p = "?a"
  输出: false
  解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。

示例 4:
  输入:
  s = "adceb"
  p = "*a*b"
  输出: true
  解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".

示例 5:
  输入:
  s = "acdcb"
  p = "a*c?b"
  输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/wildcard-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func isMatch(s string, p string) bool {
	ns, np := len(s), len(p)
	dp := make([][]bool, ns+1)
	for i := range dp {
		dp[i] = make([]bool, np+1)
	}
	dp[0][0] = true
	for i := 0; i < np && p[i] == '*'; i++ {
		dp[0][i+1] = true
	}
	for i := range s {
		for j := range p {
			if s[i] == p[j] || '?' == p[j] {
				dp[i+1][j+1] = dp[i][j]
			} else if '*' == p[j] {
				dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1]
			}
		}
	}
	return dp[ns][np]
}

// Greedy Algorithm
func isMatch2(s string, p string) bool {
	for len(s) > 0 && len(p) > 0 && p[len(p)-1] != '*' {
		if s[len(s)-1] == p[len(p)-1] || '?' == p[len(p)-1] {
			s = s[:len(s)-1]
			p = p[:len(p)-1]
		} else {
			return false
		}
	}
	if p == "" {
		return s == ""
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
