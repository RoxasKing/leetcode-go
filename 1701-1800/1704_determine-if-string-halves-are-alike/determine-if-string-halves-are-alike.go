package main

// Difficulty:
// Easy

func halvesAreAlike(s string) bool {
	isVowel := func(c rune) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}
	n := len(s)
	cnt := 0
	for _, c := range s[:n/2] {
		if isVowel(c) {
			cnt++
		}
	}
	for _, c := range s[n/2:] {
		if isVowel(c) {
			cnt--
		}
	}
	return cnt == 0
}
