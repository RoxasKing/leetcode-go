package main

// Tags:
// Two Pointers
func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		for l < r && !isValid(s[l]) {
			l++
		}
		for l < r && !isValid(s[r]) {
			r--
		}
		if lowerCase(s[l]) != lowerCase(s[r]) {
			return false
		}
	}
	return true
}

func isValid(ch byte) bool {
	return 'A' <= ch && ch <= 'Z' || 'a' <= ch && ch <= 'z' || '0' <= ch && ch <= '9'
}

func lowerCase(ch byte) byte {
	if 'A' <= ch && ch <= 'Z' {
		return ch - 'A' + 'a'
	}
	return ch
}
