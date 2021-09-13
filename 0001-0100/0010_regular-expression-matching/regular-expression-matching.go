package main

// Tags:
// Dynamic Programming

func isMatch(s string, p string) bool {
	f := [21][31]bool{{true}}
	m, n := len(s), len(p)
	for i := 1; i < n && p[i] == '*'; i += 2 {
		f[0][i+1] = true
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == p[j] || '.' == p[j] {
				f[i+1][j+1] = f[i][j]
			} else if '*' == p[j] {
				f[i+1][j+1] = f[i+1][j-1]
				if !f[i+1][j+1] && (s[i] == p[j-1] || '.' == p[j-1]) {
					f[i+1][j+1] = f[i+1][j] || f[i][j+1]
				}
			}
		}
	}
	return f[m][n]
}
