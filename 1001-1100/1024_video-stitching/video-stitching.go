package main

// Tags:
// Dynamic Programming
func videoStitching(clips [][]int, t int) int {
	dp := make([]int, t+1)
	for i := 1; i <= t; i++ {
		dp[i] = 101
	}
	for i := 1; i <= t; i++ {
		for _, clip := range clips {
			if clip[0] <= i && i <= clip[1] {
				dp[i] = Min(dp[i], dp[clip[0]]+1)
			}
		}
	}
	if dp[t] == 101 {
		return -1
	}
	return dp[t]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
