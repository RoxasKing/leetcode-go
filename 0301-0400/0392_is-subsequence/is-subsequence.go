package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	i, j := 0, 0
	for ; i < m && j < n; j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == m
}
