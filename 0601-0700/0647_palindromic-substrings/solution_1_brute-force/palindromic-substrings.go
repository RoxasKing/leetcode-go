package main

// DiffiCulty:
// Medium

// Tags:
// Brute Force

func countSubstrings(s string) int {
	o := 0
	n := len(s)
	for i := range s {
		for l, r := i, i; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
			o++
		}
		if i+1 == n || s[i] != s[i+1] {
			continue
		}
		for l, r := i, i+1; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
			o++
		}
	}
	return o
}
