package main

import "strings"

func countSegments(s string) int {
	out := 0
	strs := strings.Split(s, " ")
	for _, str := range strs {
		if str != "" {
			out++
		}
	}
	return out
}
