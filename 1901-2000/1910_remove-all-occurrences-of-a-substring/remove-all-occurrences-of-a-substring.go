package main

import "strings"

func removeOccurrences(s string, part string) string {
	for {
		i := strings.Index(s, part)
		if i == -1 {
			break
		}
		s = s[:i] + s[i+len(part):]
	}
	return s
}
