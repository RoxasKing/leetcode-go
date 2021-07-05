package main

// Tags:
// Dynamic Programming
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n, k := len(s1), len(s2), len(s3)
	if m+n != k {
		return false
	}
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i < m && dp[i][0]; i++ {
		dp[i+1][0] = s1[i] == s3[i]
	}
	for j := 0; j < n && dp[0][j]; j++ {
		dp[0][j+1] = s2[j] == s3[j]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = dp[i][j+1] && s1[i] == s3[i+j+1] || dp[i+1][j] && s2[j] == s3[i+j+1]
		}
	}
	return dp[m][n]
}

// Dynamic Programming (O(s2.length) additional memory space)
func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	m, n := len(s1), len(s2)
	dp := make([]bool, n+1)
	dp[0] = true
	for j := 0; j < n && dp[j]; j++ {
		dp[j+1] = s2[j] == s3[j]
	}
	for i := 0; i < m; i++ {
		dp[0] = dp[0] && s1[i] == s3[i]
		for j := 0; j < n; j++ {
			dp[j+1] = dp[j+1] && s1[i] == s3[i+j+1] || dp[j] && s2[j] == s3[i+j+1]
		}
	}
	return dp[n]
}
