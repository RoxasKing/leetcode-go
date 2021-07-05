package main

// Tags:
// Dynamic Programming

func strangePrinter(s string) int {
	n := len(s)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
		f[i][i] = 1
	}
	for k := 2; k <= n; k++ {
		for i := 0; i+(k-1) < n; i++ {
			f[i][i+k-1] = k
			for j := i + 1; j < i+k; j++ {
				tmp := f[i][j-1] + f[j][i+k-1]
				if s[i] == s[i+k-1] {
					tmp--
				}
				f[i][i+k-1] = Min(f[i][i+k-1], tmp)
			}
		}
	}
	return f[0][n-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
