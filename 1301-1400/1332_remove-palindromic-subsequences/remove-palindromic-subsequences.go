package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

func removePalindromeSub(s string) int {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return 2
		}
	}
	return 1
}
