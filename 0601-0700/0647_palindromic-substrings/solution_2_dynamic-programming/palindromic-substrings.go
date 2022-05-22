package main

// DiffiCulty:
// Medium

// Tags:
// Dynamic Programming

func countSubstrings(s string) int {
	n := len(s)
	f := make([]bool, n)

	out := 0
	for i := 0; i < n; i++ {
		f[i] = true
		out++
		for j := 0; j < i; j++ {
			if s[i] == s[j] && f[j+1] {
				f[j] = true
				out++
			} else {
				f[j] = false
			}
		}
	}
	return out
}
