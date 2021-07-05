package main

import "strings"

func reverseWords(s string) string {
	strs := []string{}
	for s != "" {
		for s != "" && s[0] == ' ' {
			s = s[1:]
		}
		if s == "" {
			break
		}
		i := 0
		for i < len(s) && s[i] != ' ' {
			i++
		}
		strs = append(strs, s[:i])
		s = s[i:]

	}
	for i := 0; i < len(strs)>>1; i++ {
		strs[i], strs[len(strs)-1-i] = strs[len(strs)-1-i], strs[i]
	}
	return strings.Join(strs, " ")
}
