package main

// Tags:
// Dynamic Programming

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		cnt := [2]int{}
		for k := range s {
			cnt[s[k]-'0']++
		}
		for i := m; i >= cnt[0]; i-- {
			for j := n; j >= cnt[1]; j-- {
				dp[i][j] = Max(dp[i][j], dp[i-cnt[0]][j-cnt[1]]+1)
			}
		}
	}
	return dp[m][n]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
