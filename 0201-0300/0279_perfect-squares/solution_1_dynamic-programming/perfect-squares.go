package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := range f {
		f[i] = i
		for j := 2; j*j <= i; j++ {
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
