package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func distinctSubseqII(s string) int {
	const mod int = 1e9 + 7
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1 // 0: null ; 1: not null
	last := [26]int{}
	for i := range last {
		last[i] = -1
	}
	for i, c := range s {
		idx := int(c - 'a')
		f[i+1] = f[i] * 2 % mod
		if last[idx] != -1 {
			f[i+1] = (f[i+1] - f[last[idx]] + mod) % mod
		}
		last[idx] = i
	}
	return f[n] - 1 // delete null string
}
