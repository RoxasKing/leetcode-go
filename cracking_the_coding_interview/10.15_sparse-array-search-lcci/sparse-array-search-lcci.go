package main

// Tags:
// Binary Search

func findString(words []string, s string) int {
	n := len(words)
	l, r := 0, n-1
	for l < r {
		m := (l + r) >> 1
		for l < m && words[m] == "" {
			m--
		}
		if words[m] < s {
			l = m + 1
		} else {
			r = m
		}
	}
	if words[l] == s {
		return l
	}
	return -1
}
