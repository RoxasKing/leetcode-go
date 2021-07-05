package main

// Tags:
// Dynamic Programming
func kInversePairs(n int, k int) int {
	mod := int(1e9 + 7)
	dp := make([]int, k+1)
	dp[0] = 1
	for num := 1; num <= n; num++ {
		dpNew := make([]int, k+1)
		dpNew[0] = 1
		for pair := 1; pair <= k && pair <= num*(num-1)>>1; pair++ {
			dpNew[pair] = dpNew[pair-1] + dp[pair]
			if pair-num >= 0 {
				dpNew[pair] -= dp[pair-num]
			}
			dpNew[pair] = (dpNew[pair] + mod) % mod
		}
		dp = dpNew
	}
	return dp[k]
}
