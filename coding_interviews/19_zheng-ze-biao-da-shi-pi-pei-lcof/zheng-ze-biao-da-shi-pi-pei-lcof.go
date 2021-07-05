package main

// Tags:
// Dynamic Programming
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j < n && dp[0][j-1]; j += 2 { // if s = "xyz", p = "a*b*xyz"
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == p[j] || '.' == p[j] {
				dp[i+1][j+1] = dp[i][j]
			} else if '*' == p[j] {
				dp[i+1][j+1] = dp[i+1][j-1] // appear 0 times
				if s[i] == p[j-1] || '.' == p[j-1] {
					dp[i+1][j+1] = dp[i+1][j+1] || dp[i][j+1] || dp[i+1][j] // appear 1 or more times
				}
			}
		}
	}
	return dp[m][n]
}
