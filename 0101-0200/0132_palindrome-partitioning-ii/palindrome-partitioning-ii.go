package main

// Tags:
// Dynamic Programming
func minCut(s string) int {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for r := range isPalindrome {
		isPalindrome[r] = make([]bool, n)
		for l := 0; l <= r; l++ {
			if s[l] == s[r] && (r-l < 2 || isPalindrome[l+1][r-1]) {
				isPalindrome[l][r] = true
			}
		}
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = i
	}

	for r := 1; r < n; r++ {
		if isPalindrome[0][r] {
			dp[r] = 0
			continue
		}
		for l := 1; l <= r; l++ {
			if isPalindrome[l][r] {
				dp[r] = Min(dp[r], dp[l-1]+1)
			}
		}
	}

	return dp[n-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
