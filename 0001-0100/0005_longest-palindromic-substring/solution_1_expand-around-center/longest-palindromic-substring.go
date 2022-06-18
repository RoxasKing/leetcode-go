package main

// Difficulty:
// Medium

func longestPalindrome(s string) string {
	out := ""
	for i := range s {
		if res := getlongetPalindrome(s, i, i); len(res) > len(out) {
			out = res
		}
		if i+1 == len(s) || s[i+1] != s[i] {
			continue
		}
		if res := getlongetPalindrome(s, i, i+1); len(res) > len(out) {
			out = res
		}
	}
	return out
}

func getlongetPalindrome(s string, l, r int) string {
	for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
		l--
		r++
	}
	return s[l : r+1]
}
