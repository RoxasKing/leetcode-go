package main

// Difficulty:
// Easy

func truncateSentence(s string, k int) string {
	i := 0
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			k--
		}
		if k == 0 {
			break
		}
	}
	return s[:i]
}
