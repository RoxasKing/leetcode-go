package main

import "strings"

// Difficulty:
// Easy

func toGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")
	t := "a"
	for i, word := range words {
		if !isVowel(word[0]) {
			words[i] = word[1:] + word[:1]
		}
		words[i] += "ma" + t
		t += "a"
	}
	return strings.Join(words, " ")
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
