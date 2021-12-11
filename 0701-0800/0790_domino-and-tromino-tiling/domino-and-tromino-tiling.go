package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func numTilings(n int) int {
	if n < 3 {
		return n
	}
	f := make([]int, n+1)
	p := make([]int, n+1)
	f[1] = 1
	f[2] = 2
	p[2] = 1
	for i := 3; i <= n; i++ {
		f[i] = (f[i-1] + f[i-2] + p[i-1]<<1) % MOD
		p[i] = (p[i-1] + f[i-2]) % MOD
	}
	return f[n]
}

const MOD int = 1e9 + 7
