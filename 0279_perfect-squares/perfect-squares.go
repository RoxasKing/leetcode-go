package main

/*
  给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
*/

// Dynamic Programming
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = i
	}
	squareNums := getSquareNums(n)
	for num := 1; num <= n; num++ {
		for i := 0; i < len(squareNums) && squareNums[i] <= num; i++ {
			dp[num] = Min(dp[num], dp[num-squareNums[i]]+1)
		}
	}
	return dp[n]
}

func getSquareNums(limit int) []int {
	l, r := 0, limit
	for l < r {
		m := l + (r-l)>>1
		sqrs := m * m
		if sqrs < limit {
			l = m + 1
		} else {
			r = m
		}
	}
	out := make([]int, l)
	for i := 0; i < l; i++ {
		out[i] = (i + 1) * (i + 1)
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
