package main

// Tags:
// Dynamic Programming
func minJump(jump []int) int {
	n := len(jump)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = n
	}
	dp[0] = 0

	out := 1<<31 - 1
	for i := range dp {
		j := i + jump[i]
		if j >= n {
			out = Min(out, dp[i]+1)
			continue
		}
		if dp[j] <= dp[i]+1 {
			continue
		}
		dp[j] = dp[i] + 1
		for k := j - 1; dp[k] >= dp[j]+1; k-- {
			dp[k] = dp[j] + 1
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Dynamic Programming
func minJump2(jump []int) int {
	n := len(jump)
	maxDist := make([]int, n)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = n
	}

	out := n
	times := 0
	for i := 0; i < n; i++ {
		if i > maxDist[times] {
			times++
		}
		dp[i] = Min(dp[i], times+1)

		next := i + jump[i]
		if next >= n {
			out = Min(out, dp[i]+1)
		} else {
			maxDist[dp[i]+1] = Max(maxDist[dp[i]+1], next)
			dp[next] = Min(dp[next], dp[i]+1)
		}
	}
	return out
}
