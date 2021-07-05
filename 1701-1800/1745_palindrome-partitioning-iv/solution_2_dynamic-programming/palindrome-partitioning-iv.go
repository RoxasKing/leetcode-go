package main

// Tags:
// Dynamic Programming

func checkPartitioning(s string) bool {
	n := len(s)
	f := make([][]bool, n)
	for r := 0; r < n; r++ {
		f[r] = make([]bool, n)
		for l := 0; l <= r; l++ {
			if s[l] == s[r] && (r-l < 2 || f[l+1][r-1]) {
				f[l][r] = true
			}
		}
	}

	for i := 1; i < n-1; i++ {
		for j := i; j < n-1; j++ {
			if f[0][i-1] && f[i][j] && f[j+1][n-1] {
				return true
			}
		}
	}
	return false
}
