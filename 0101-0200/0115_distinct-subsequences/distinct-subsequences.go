package main

// Tags:
// Dynamic Programming

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		f[i][0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			f[i+1][j+1] = f[i][j+1]
			if s[i] == t[j] {
				f[i+1][j+1] += f[i][j]
			}
		}
	}
	return f[m][n]
}
