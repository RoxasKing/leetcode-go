package main

// Difficulty:
// Easy

func minOperations(s string) int {
	a, b := 0, 0
	for i, c := range s {
		if i&1 == 0 && c == '1' || i&1 == 1 && c == '0' {
			a++
		}
		if i&1 == 0 && c == '0' || i&1 == 1 && c == '1' {
			b++
		}
	}
	if a < b {
		return a
	}
	return b
}
