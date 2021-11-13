package main

// Difficulty:
// Easy

func detectCapitalUse(word string) bool {
	a, b := 0, 0
	for i := range word {
		if 'a' <= word[i] && word[i] <= 'z' {
			a++
		} else {
			b++
		}
	}
	return a == 0 || b == 0 || b == 1 && word[0] < 'a'
}
