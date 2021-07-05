package main

// Tags:
// Math + Dynamic Programming
func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1
	}
	dp := make([]float64, K+W)
	for i := K; i <= N && i <= K-1+W; i++ {
		dp[i] = 1
	}
	dp[K-1] = float64(Min(N, K-1+W)-(K-1)) / float64(W)
	for i := K - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - (dp[i+W+1]-dp[i+1])/float64(W)
	}
	return dp[0]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
