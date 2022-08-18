package main

// Difficulty:
// Easy

// Tags:
// Hash

func firstUniqChar(s string) int {
	f := [26]int{}
	for _, c := range s {
		f[c-'a']++
	}
	for i, c := range s {
		if f[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
