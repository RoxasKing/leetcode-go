package main

// Tags:
// Dynamic Programming
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	f := make([][]bool, m+1)
	for i := range f {
		f[i] = make([]bool, n+1)
	}

	f[0][0] = true
	for i := 1; i < n && f[0][i-1]; i += 2 {
		if p[i] == '*' {
			f[0][i+1] = true
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == p[j] || '.' == p[j] {
				f[i+1][j+1] = f[i][j]
			} else if '*' == p[j] {
				f[i+1][j+1] = f[i+1][j-1] // repeat 0 times
				if s[i] == p[j-1] || '.' == p[j-1] {
					f[i+1][j+1] = f[i+1][j+1] || f[i+1][j] || f[i][j+1] // repeat 1 or more times
				}
			}
		}
	}
	return f[m][n]
}
