package main

/*
  硬币。给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。(结果可能会很大，你需要将结果模上1000000007)
*/

func waysToChange(n int) int {
	mod := 1000000007
	coins := []int{25, 10, 5, 1}
	dp := make([]int, n+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			dp[i] = (dp[i] + dp[i-coin]) % mod
		}
	}
	return dp[n]
}

func waysToChange2(n int) int {
	var out int
	mod := 1000000007
	for i := 0; i*25 <= n; i++ {
		rest := n - i*25
		a, b := rest/10, rest%10/5
		out = (out + (a+1)*(a+b+1)%mod) % mod
	}
	return out
}
