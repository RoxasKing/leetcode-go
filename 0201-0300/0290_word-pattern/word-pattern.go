package main

import "strings"

// Difficulty:
// Easy

// Tags:
// Hash

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(strs) != len(pattern) {
		return false
	}
	mark := map[string]byte{}
	dict := [26]string{}
	for i, str := range strs {
		x := pattern[i] - 'a'
		if y, ok := mark[str]; !ok {
			mark[str] = x
		} else if y != x {
			return false
		}
		if dict[x] == "" {
			dict[x] = str
		} else if dict[x] != str {
			return false
		}
	}
	return true
}
