package main

// Difficulty:
// Easy

// Tags:
// Hash

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	h := [26]int{}
	for _, c := range s {
		h[c-'a']++
	}
	for _, c := range t {
		if h[c-'a']--; h[c-'a'] < 0 {
			return false
		}
	}
	return true
}
