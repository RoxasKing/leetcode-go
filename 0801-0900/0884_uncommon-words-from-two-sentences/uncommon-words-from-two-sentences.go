package main

import "strings"

// Difficulty:
// Easy

// Tags:
// Hash

func uncommonFromSentences(s1 string, s2 string) []string {
	freq := map[string]int{}
	for _, str := range strings.Split(s1, " ") {
		freq[str]++
	}
	for _, str := range strings.Split(s2, " ") {
		freq[str]++
	}
	out := []string{}
	for str, cnt := range freq {
		if cnt == 1 {
			out = append(out, str)
		}
	}
	return out
}
