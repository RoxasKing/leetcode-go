package main

// Tags:
// Dynamic Programming

func getMoneyAmount(n int) int {
	f := make([][]int, n+2)
	for i := range f {
		f[i] = make([]int, n+2)
	}

	for i := 1; i < n; i++ {
		for j := 1; i+j <= n; j++ {
			f[j][j+i] = 1<<31 - 1
			for k := j; k <= i+j; k++ {
				f[j][j+i] = Min(f[j][j+i], k+Max(f[j][k-1], f[k+1][j+i]))
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
