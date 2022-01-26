package main

// DIfficulty:
// Easy

// Tags:
// Dynamic Programming

func numberOfMatches(n int) int {
	f := make([]int, n+1)
	for i := 2; i <= n; i++ {
		f[i] = f[i>>1+i&1] + i>>1
	}
	return f[n]
}
