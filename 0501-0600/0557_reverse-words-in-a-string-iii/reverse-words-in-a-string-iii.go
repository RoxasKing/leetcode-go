package main

import "strings"

// Difficulty:
// Easy

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	for i, str := range strs {
		a := []byte(str)
		for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
			a[l], a[r] = a[r], a[l]
		}
		strs[i] = string(a)
	}
	return strings.Join(strs, " ")
}
