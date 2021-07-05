package main

// Tags:
// Two Pointers
func countBinarySubstrings(s string) int {
	out := 0
	n := len(s)
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			continue
		}
		for l, r := i, i+1; l >= 0 && s[l] == s[i] && r < n && s[r] == s[i+1]; l, r = l-1, r+1 {
			out++
		}
	}
	return out
}
