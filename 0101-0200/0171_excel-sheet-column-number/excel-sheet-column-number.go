package main

// Difficulty:
// Easy

func titleToNumber(columnTitle string) int {
	out := 0
	for base, i := 1, len(columnTitle)-1; i >= 0; base, i = 26*base, i-1 {
		out += int(columnTitle[i]-'A'+1) * base
	}
	return out
}
