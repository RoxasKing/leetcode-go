package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) == 0 && len(s2) == 0 {
		return true
	} else if len(s1) != len(s2) {
		return false
	}
	for i, n := 0, len(s1); i < n; i++ {
		ok := true
		for j := 0; j < n; j++ {
			if s1[(i+j)%n] != s2[j] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}
