package main

import "strings"

// Difficulty:
// Medium

// Tags:
// Two Pointers

func reverseWords(s string) string {
	a := strings.Split(s, " ")
	strs := []string{}
	for _, s := range a {
		if len(s) > 0 {
			strs = append(strs, s)
		}
	}
	for l, r := 0, len(strs)-1; l < r; l, r = l+1, r-1 {
		strs[l], strs[r] = strs[r], strs[l]
	}
	return strings.Join(strs, " ")
}
