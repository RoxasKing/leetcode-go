package main

import "strings"

// Difficulty:
// Easy

func findOcurrences(text, first, second string) []string {
	words := strings.Split(text, " ")
	var out []string
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			out = append(out, words[i])
		}
	}
	return out
}
