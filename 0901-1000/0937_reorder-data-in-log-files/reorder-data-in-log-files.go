package main

import (
	"sort"
)

// Difficulty:
// Easy

// Tags:
// Sorting

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		a, b := logs[i], logs[j]
		m, n := 0, 0
		for ; a[m] != ' '; m++ {
		}
		for ; b[n] != ' '; n++ {
		}
		m++
		n++
		isDigA := '0' <= a[m] && a[m] <= '9'
		isDigB := '0' <= b[n] && b[n] <= '9'
		if isDigA && isDigB {
			return false
		}
		if !isDigA && !isDigB {
			c, d := a[m:], b[n:]
			return c < d || c == d && a < b
		}
		return !isDigA
	})
	return logs
}
