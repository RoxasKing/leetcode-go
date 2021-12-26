package main

// Difficulty:
// Easy

func toLowerCase(s string) string {
	chs := make([]byte, len(s))
	for i := range s {
		chs[i] = s[i]
		if 'A' <= s[i] && s[i] <= 'Z' {
			chs[i] += 'a' - 'A'
		}
	}
	return string(chs)
}
