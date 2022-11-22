package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func soupServings(n int) float64 {
	if n = (n + 24) / 25; n >= 179 {
		return 1
	}
	f := make([][]float64, n+1)
	for i := range f {
		f[i] = make([]float64, n+1)
	}
	f[0][0] = 0.5
	for i := 1; i <= n; i++ {
		f[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			f[i][j] = (f[max(0, i-4)][j] +
				f[max(0, i-3)][max(0, j-1)] +
				f[max(0, i-2)][max(0, j-2)] +
				f[max(0, i-1)][max(0, j-3)]) / 4
		}
	}
	return f[n][n]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
