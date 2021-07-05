package main

// Tags:
// Dynamic Programming

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, target+1)
			for k := 1; k <= target; k++ {
				dp[i][j][k] = 1<<31 - 1
			}
		}
	}

	if houses[0] == 0 {
		for j := 0; j < n; j++ {
			dp[0][j][1] = cost[0][j]
		}
	} else {
		dp[0][houses[0]-1][1] = 0
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 1; k <= target; k++ {
				if dp[i-1][j][k] == 1<<31-1 {
					continue
				}
				if houses[i] > 0 {
					if j == houses[i]-1 {
						dp[i][houses[i]-1][k] = Min(dp[i][houses[i]-1][k], dp[i-1][j][k])
					} else if k < target {
						dp[i][houses[i]-1][k+1] = Min(dp[i][houses[i]-1][k+1], dp[i-1][j][k])
					}
					continue
				}
				for nj := 0; nj < n; nj++ {
					if nj == j {
						dp[i][nj][k] = Min(dp[i][nj][k], dp[i-1][j][k]+cost[i][nj])
					} else if k < target {
						dp[i][nj][k+1] = Min(dp[i][nj][k+1], dp[i-1][j][k]+cost[i][nj])
					}
				}
			}
		}
	}

	out := 1<<31 - 1
	for j := 0; j < n; j++ {
		out = Min(out, dp[m-1][j][target])
	}

	if out == 1<<31-1 {
		return -1
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
