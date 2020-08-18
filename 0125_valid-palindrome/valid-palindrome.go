package main

/*
  给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
  说明：本题中，我们将空字符串定义为有效的回文串。
*/

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		for l < r && !isValid(s[l]) {
			l++
		}
		for l < r && !isValid(s[r]) {
			r--
		}
		cl, cr := s[l], s[r]
		lowerCase(&cl)
		lowerCase(&cr)
		if cl != cr {
			return false
		}
	}
	return true
}

func isValid(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9'
}

func lowerCase(c *byte) {
	if *c < 'a' {
		*c = *c - 'A' + 'a'
	}
}
