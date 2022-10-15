package main

// Difficulty:
// Easy

// Tags:
// Hash

func areAlmostEqual(s1 string, s2 string) bool {
	h := [26]int{}
	for _, c := range s1 {
		h[c-'a']++
	}
	cnt := 0
	for i, c := range s2 {
		if h[c-'a']--; h[c-'a'] < 0 {
			return false
		}
		if s1[i] != s2[i] {
			if cnt++; cnt > 2 {
				return false
			}
		}
	}
	return cnt == 0 || cnt == 2
}
