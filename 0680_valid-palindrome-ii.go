package main

/*
  给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
  注意:
    字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
*/

func validPalindrome(s string) bool {
	isPalindrome := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isPalindrome(l+1, r) || isPalindrome(l, r-1)
		}
		l++
		r--
	}
	return true
}
