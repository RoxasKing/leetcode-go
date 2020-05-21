package leetcode

/*
  给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
*/

func longestPalindrome(s string) string {
	var out string
	for i := range s {
		out = maxStr(out, palindrome(s, i, i))
		if i+1 < len(s) && s[i] == s[i+1] {
			out = maxStr(out, palindrome(s, i, i+1))
		}
	}
	return out
}

func palindrome(s string, l, r int) string {
	for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
		l--
		r++
	}
	return s[l : r+1]
}

func maxStr(s1, s2 string) string {
	if len(s1) < len(s2) {
		return s2
	}
	return s1
}
