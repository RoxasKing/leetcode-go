package main

// Tags:
// DFS
func minSteps(n int) int {
	out := n
	dfs(n, 1, 0, 0, &out)
	return out
}

func dfs(n, cur, steps, prevStep int, out *int) {
	if steps >= *out {
		return
	}

	if cur == n {
		*out = Min(*out, steps)
		return
	}

	if cur+cur <= n {
		dfs(n, cur+cur, steps+2, cur, out)
	}

	if prevStep > 0 && cur+prevStep <= n {
		dfs(n, cur+prevStep, steps+1, prevStep, out)
	}
}

// Dynamic Programming
func minSteps2(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = 1<<31 - 1
		for j := 1; j < i; j++ {
			if i%j == 0 {
				dp[i] = Min(dp[i], dp[j]+i/j)
			}
		}
	}
	return dp[n]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
