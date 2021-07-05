package main

// Tags:
// Dynamic Programming
func numPermsDISequence(S string) int {
	mod := int(1e9 + 7)
	n := len(S)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			l, r := 0, j // S[i-1] == 'I'
			if S[i-1] == 'D' {
				l, r = j, i
			}
			for k := l; k < r; k++ {
				dp[i][j] += dp[i-1][k]
				dp[i][j] %= mod
			}
		}
	}

	out := 0
	for _, cnt := range dp[n] {
		out += cnt
		out %= mod
	}
	return out
}

// Dynamic Programming (O(S.length) additional memory space)
func numPermsDISequence2(S string) int {
	mod := int(1e9 + 7)
	n := len(S)
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			l, r := 0, j // S[i-1] == 'I'
			if S[i-1] == 'D' {
				l, r = j, i
			}
			dp[1][j] = 0
			for k := l; k < r; k++ {
				dp[1][j] += dp[0][k]
				dp[1][j] %= mod
			}
		}
		dp[0], dp[1] = dp[1], dp[0]
	}

	out := 0
	for j := range dp[0] {
		out += dp[0][j]
		out %= mod
	}
	return out
}
