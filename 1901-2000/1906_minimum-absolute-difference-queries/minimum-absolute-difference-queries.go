package main

// Tags:
// Prefix Sum

func minDifference(nums []int, queries [][]int) []int {
	m, n := len(nums), len(queries)
	for i := 1; i <= m; i++ {
		for j := 1; j <= 100; j++ {
			cnt[i][j] = cnt[i-1][j]
		}
		cnt[i][nums[i-1]]++
	}
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = -1
		l, r := queries[i][0], queries[i][1]
		pre, ans := -int(1e9), int(1e9)
		for j := 1; j <= 100; j++ {
			if cnt[r+1][j] != cnt[l][j] {
				ans = Min(ans, j-pre)
				pre = j
			}
		}
		if ans != int(1e9) {
			out[i] = ans
		}
	}
	return out
}

var cnt = [1e5 + 1][101]int{}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
