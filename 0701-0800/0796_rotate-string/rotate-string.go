package main

// Difficulty:
// Easy

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i:]+s[:i] == goal {
			return true
		}
	}
	return false
}
