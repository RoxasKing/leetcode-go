package main

func countGoodSubstrings(s string) int {
	n := len(s)
	out := 0
	for i := 0; i < n-2; i++ {
		if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
			out++
		}
	}
	return out
}
