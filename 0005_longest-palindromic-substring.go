package My_LeetCode_In_Go

/*
  给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
*/

func longestPalindrome(s string) string {
	var out string
	for i := range s {
		out = maxStr(out, maxStr(palindrome(s, i, i), palindrome(s, i, i+1)))
	}
	return out
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l, r = l-1, r+1
	}
	return s[l+1 : r]
}

func maxStr(s1, s2 string) string {
	if len(s1) < len(s2) {
		return s2
	}
	return s1
}