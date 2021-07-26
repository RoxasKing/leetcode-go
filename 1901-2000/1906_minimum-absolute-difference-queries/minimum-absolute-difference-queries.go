package main

// Tags:
// Prefix Sum

func minDifference(nums []int, queries [][]int) []int {
	m, n := len(nums), len(queries)
	cnt := make([][101]int16, m+1)
	for i := 1; i <= m; i++ {
		for j := 1; j <= 100; j++ {
			cnt[i][j] = cnt[i-1][j]
		}
		cnt[i][nums[i-1]]++
	}
	out := make([]int, n)
	for i, q := range queries {
		out[i] = -1
		l, r, pre := q[0], q[1], 0
		for j := 1; j <= 100; j++ {
			if cnt[r+1][j] != cnt[l][j] {
				if pre > 0 && (out[i] == -1 || j-pre < out[i]) {
					out[i] = j - pre
				}
				pre = j
			}
		}
	}
	return out
}
