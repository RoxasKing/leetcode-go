package main

// Difficulty:
// Easy

// Tags:
// Hash

func CheckPermutation(s1 string, s2 string) bool {
	h := [26]int{}
	for _, c := range s1 {
		h[c-'a']++
	}
	for _, c := range s2 {
		if h[c-'a']--; h[c-'a'] < 0 {
			return false
		}
	}
	return true
}
