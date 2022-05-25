package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func findMaxForm(strs []string, m int, n int) int {
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for _, s := range strs {
		a, b := 0, 0
		for _, c := range s {
			if c == '0' {
				a++
			} else {
				b++
			}
		}
		for i := m; i >= a; i-- {
			for j := n; j >= b; j-- {
				f[i][j] = max(f[i][j], f[i-a][j-b]+1)
			}
		}
	}
	return f[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
