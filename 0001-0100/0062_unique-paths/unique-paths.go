package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func uniquePaths(m int, n int) int {
	f := make([]int, n)
	f[0] = 1
	for ; m > 0; m-- {
		for j := 1; j < n; j++ {
			f[j] += f[j-1]
		}
	}
	return f[n-1]
}
