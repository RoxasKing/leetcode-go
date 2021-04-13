package main

/*
  编写一个函数来查找字符串数组中的最长公共前缀。
  如果不存在公共前缀，返回空字符串 ""。

  示例 1:
    输入: ["flower","flow","flight"]
    输出: "fl"

  示例 2:
    输入: ["dog","racecar","car"]
    输出: ""
    解释: 输入不存在公共前缀。

  说明:
    所有输入只包含小写字母 a-z 。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-common-prefix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for j := range strs[0] {
		for i := 1; i < len(strs); i++ {
			if j == len(strs[i]) || strs[i][j] != strs[0][j] {
				return strs[0][:j]
			}
		}
	}
	return strs[0]
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	lcp := func(str1, str2 string) string {
		if len(str1) > len(str2) {
			str1, str2 = str2, str1
		}
		for i := range str1 {
			if str1[i] != str2[i] {
				return str1[:i]
			}
		}
		return str1
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
	}
	return prefix
}

// Divide-and-conquer
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(l, r int) string {
		if l == r {
			return strs[l]
		}
		m := l + (r-l)>>1
		lPrefix, rPrefix := lcp(l, m), lcp(m+1, r)
		if len(lPrefix) > len(rPrefix) {
			lPrefix, rPrefix = rPrefix, lPrefix
		}
		for i := range lPrefix {
			if lPrefix[i] != rPrefix[i] {
				return lPrefix[:i]
			}
		}
		return lPrefix
	}
	return lcp(0, len(strs)-1)
}

// Binary Search
func longestCommonPrefix4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	isCommonPrefix := func(r int) bool {
		for i := 1; i < len(strs); i++ {
			if strs[0][:r] != strs[i][:r] {
				return false
			}
		}
		return true
	}
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		minLen = Min(minLen, len(strs[i]))
	}
	l, r := 0, minLen-1
	for l <= r {
		m := l + (r-l)>>1
		if isCommonPrefix(m + 1) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return strs[0][:l]
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
