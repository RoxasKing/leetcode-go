package main

// Tags:
// Dynamic Programming
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	out := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			}
			out = Max(out, dp[i+1][j+1])
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Sliding Window
func findLength2(nums1 []int, nums2 []int) int {
	out := 0
	m, n := len(nums1), len(nums2)
	for i := 0; i < m && m-i > out; i++ {
		out = Max(out, maxLen(nums1[i:], nums2))
	}
	for i := 0; i < n && n-i > out; i++ {
		out = Max(out, maxLen(nums1, nums2[i:]))
	}
	return out
}

func maxLen(nums1, nums2 []int) int {
	out := 0
	for i, cnt := 0, 0; i < len(nums1) && i < len(nums2); i++ {
		if nums1[i] != nums2[i] {
			cnt = 0
		} else {
			cnt++
			out = Max(out, cnt)
		}
	}
	return out
}
