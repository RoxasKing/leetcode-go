package main

import "strings"

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(strs) != len(pattern) {
		return false
	}
	mark := make(map[string]byte)
	dict := [26]string{}
	for i := range pattern {
		if val, ok := mark[strs[i]]; !ok {
			mark[strs[i]] = pattern[i]
		} else if val != pattern[i] {
			return false
		}
		if dict[pattern[i]-'a'] == "" {
			dict[pattern[i]-'a'] = strs[i]
		} else if dict[pattern[i]-'a'] != strs[i] {
			return false
		}
	}
	return true
}
