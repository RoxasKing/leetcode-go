package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func numTrees(n int) int {
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			f[i] += f[j] * f[i-1-j]
		}
	}
	return f[n]
}
