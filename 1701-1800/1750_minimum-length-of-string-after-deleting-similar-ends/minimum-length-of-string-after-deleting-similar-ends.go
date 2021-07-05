package main

// Tags:
// Two Pointers

func minimumLength(s string) int {
	n := len(s)
	l, r := 0, n-1
	for l < r && s[l] == s[r] {
		c := s[l]
		for l <= r && s[l] == c {
			l++
		}
		for l < r && s[r] == c {
			r--
		}
	}
	return r + 1 - l
}
