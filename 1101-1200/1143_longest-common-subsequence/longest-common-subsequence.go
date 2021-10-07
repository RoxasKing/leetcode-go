package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = Max(f[i+1][j], f[i][j+1])
			}
		}
	}
	return f[m][n]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
