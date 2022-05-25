package main

// DiffiCulty:
// Medium

// Tags:
// Dynamic Programming

func countSubstrings(s string) int {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		f[i][i] = true
	}
	o := n
	for k := 0; k < n; k++ {
		for i := 0; i+k < n; i++ {
			if (i+1 > i+k-1 || f[i+1][i+k-1]) && s[i] == s[i+k] {
				f[i][i+k] = true
				o++
			}
		}
	}
	return o
}
