package main

// Difficulty:
// Easy

func maxDepth(s string) int {
	var out, depth int
	for i := range s {
		if s[i] == '(' {
			depth++
			if depth > out {
				out = depth
			}
		} else if s[i] == ')' {
			depth--
		}
	}
	return out
}
