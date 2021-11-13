package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func getMoneyAmount(n int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for l := 1; l < n; l++ {
		for i := 1; i+l <= n; i++ {
			f[i][i+l] = 1<<31 - 1
			for k := i; k < i+l; k++ {
				f[i][i+l] = Min(f[i][i+l], k+Max(f[i][k-1], f[k+1][i+l]))
			}
		}
	}
	return f[1][n]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
