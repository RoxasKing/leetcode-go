package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

func reverseOnlyLetters(s string) string {
	chs := []byte(s)
	for l, r := 0, len(s)-1; ; l, r = l+1, r-1 {
		for ; l < len(s) && !isLetter(s[l]); l++ {
		}
		for ; r >= 0 && !isLetter(s[r]); r-- {
		}
		if l >= r {
			break
		}
		chs[l], chs[r] = chs[r], chs[l]
	}
	return string(chs)
}

func isLetter(ch byte) bool { return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' }
