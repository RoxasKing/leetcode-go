package main

/*
  给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
  返回符合要求的最少分割次数。
  示例:
    输入: "aab"
    输出: 1
    解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
*/

func minCut(s string) int {
	check := make([][]bool, len(s))
	for r := range check {
		check[r] = make([]bool, len(s))
		for l := 0; l <= r; l++ {
			if s[l] == s[r] && (r-l <= 2 || check[l+1][r-1]) {
				check[l][r] = true
			}
		}
	}
	dp := make([]int, len(s))
	for i := range dp {
		dp[i] = i
	}
	for i := 1; i < len(s); i++ {
		if check[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if check[j+1][i] && dp[i] > dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[len(s)-1]
}
