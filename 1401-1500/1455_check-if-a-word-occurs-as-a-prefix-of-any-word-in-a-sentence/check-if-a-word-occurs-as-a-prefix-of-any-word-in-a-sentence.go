package main

import "strings"

// Difficulty:
// Easy

func isPrefixOfWord(sentence string, searchWord string) int {
	ws := strings.Split(sentence, " ")
	for i, w := range ws {
		if len(w) >= len(searchWord) && w[:len(searchWord)] == searchWord {
			return i + 1
		}
	}
	return -1
}
