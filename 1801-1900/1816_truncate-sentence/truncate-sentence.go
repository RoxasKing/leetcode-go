package main

import "strings"

func truncateSentence(s string, k int) string {
	ss := strings.Split(s, " ")
	if len(ss) <= k {
		return s
	}
	return strings.Join(ss[:k], " ")
}
