package main

// Difficulty:
// Easy

func detectCapitalUse(word string) bool {
	cnt := 0
	for i := range word {
		if 'A' <= word[i] && word[i] <= 'Z' {
			cnt++
		}
	}
	return cnt == 0 || cnt == len(word) || cnt == 1 && 'A' <= word[0] && word[0] <= 'Z'
}
