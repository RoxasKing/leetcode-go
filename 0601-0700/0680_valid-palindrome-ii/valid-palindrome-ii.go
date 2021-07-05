package main

// Tags:
// Two Pointers
func validPalindrome(s string) bool {
	n := len(s)
	l, r := 0, n-1
	for l < r {
		if s[l] != s[r] {
			return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
