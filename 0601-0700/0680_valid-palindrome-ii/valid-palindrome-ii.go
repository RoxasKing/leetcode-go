package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for ; l < r && s[l] == s[r]; l, r = l+1, r-1 {
	}
	return l >= r || valid(s[l+1:r+1]) || valid(s[l:r])
}

func valid(s string) bool {
	l, r := 0, len(s)-1
	for ; l < r && s[l] == s[r]; l, r = l+1, r-1 {
	}
	return l >= r
}
