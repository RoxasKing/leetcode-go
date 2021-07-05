package main

// Tags:
// Dynamic Programming
func maxDotProduct(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
		for j := range dp[i] {
			dp[i][j] = -1 << 31
		}
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			mul := nums1[i-1] * nums2[j-1]
			dp[i][j] = Max(dp[i][j], Max(Max(dp[i-1][j-1]+mul, mul), Max(dp[i-1][j], dp[i][j-1])))
		}
	}
	return dp[n1][n2]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
