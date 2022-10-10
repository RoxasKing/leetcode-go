package main

// Difficulty:
// Easy

func checkOnesSegment(s string) bool {
	for ; len(s) > 0 && s[0] == '1'; s = s[1:] {
	}
	for ; len(s) > 0 && s[0] == '0'; s = s[1:] {
	}
	return len(s) == 0
}
