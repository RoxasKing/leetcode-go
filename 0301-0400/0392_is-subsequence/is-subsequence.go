package main

// Tags:
// Two Pointers
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	i := 0
	for j := 0; i < m && j < n; j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == m
}
