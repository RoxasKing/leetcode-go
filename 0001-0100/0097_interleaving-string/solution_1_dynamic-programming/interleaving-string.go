package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	f := make([][]bool, m+1)
	for i := range f {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i < m && f[i][0]; i++ {
		f[i+1][0] = s1[i] == s3[i]
	}
	for j := 0; j < n && f[0][j]; j++ {
		f[0][j+1] = s2[j] == s3[j]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			f[i+1][j+1] = f[i][j+1] && s1[i] == s3[i+j+1] || f[i+1][j] && s2[j] == s3[i+j+1]
		}
	}
	return f[m][n]
}
