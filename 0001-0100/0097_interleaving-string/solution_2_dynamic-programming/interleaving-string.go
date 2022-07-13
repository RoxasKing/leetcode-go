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
	f := make([]bool, n+1)
	f[0] = true
	for j := 0; j < n && f[j]; j++ {
		f[j+1] = s2[j] == s3[j]
	}
	for i := 0; i < m; i++ {
		f[0] = f[0] && s1[i] == s3[i]
		for j := 0; j < n; j++ {
			f[j+1] = f[j+1] && s1[i] == s3[i+j+1] || f[j] && s2[j] == s3[i+j+1]
		}
	}
	return f[n]
}
