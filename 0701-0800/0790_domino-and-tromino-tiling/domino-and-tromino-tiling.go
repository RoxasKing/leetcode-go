package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func numTilings(n int) int {
	if n < 2 {
		return n
	}
	const mod int = 1e9 + 7
	f, p := make([]int, n+1), make([]int, n+1)
	f[1], f[2], p[2] = 1, 2, 1
	for i := 3; i <= n; i++ {
		f[i] = (f[i-1] + f[i-2] + p[i-1]<<1) % mod
		p[i] = (p[i-1] + f[i-2]) % mod
	}
	return f[n]
}
