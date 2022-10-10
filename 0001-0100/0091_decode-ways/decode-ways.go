package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func numDecodings(s string) int {
	f0, f1 := 0, 1
	for i := 0; i < len(s); i++ {
		nf := 0
		if i > 0 && (s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6') {
			nf += f0
		}
		if s[i] != '0' {
			nf += f1
		}
		f0, f1 = f1, nf
	}
	return f1
}
