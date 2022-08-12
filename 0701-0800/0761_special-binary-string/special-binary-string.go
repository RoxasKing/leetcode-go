package main

import (
	"sort"
	"strings"
)

// Difficulty:
// Hard

// Tags:
// Recursion

func makeLargestSpecial(s string) string {
	a := []string{}
	for i, j, c := 0, 0, 0; j < len(s); j++ {
		if s[j] == '1' {
			c++
		} else {
			c--
		}
		if c == 0 {
			a = append(a, "1"+makeLargestSpecial(s[i+1:j])+"0")
			i = j + 1
		}
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	return strings.Join(a, "")
}
