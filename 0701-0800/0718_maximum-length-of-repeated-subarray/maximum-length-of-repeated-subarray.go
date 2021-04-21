package main

/*
  Given two integer arrays nums1 and nums2, return the maximum length of a subarray that appears in both arrays.

  Example 1:
    Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
    Output: 3
    Explanation: The repeated subarray with maximum length is [3,2,1].

  Example 2:
    Input: nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
    Output: 5

  Constraints:
    1. 1 <= nums1.length, nums2.length <= 1000
    2. 0 <= nums1[i], nums2[i] <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func findLength(A []int, B []int) int {
	out := 0
	m, n := len(A), len(B)
	for i := 0; i < m && m-i > out; i++ {
		out = Max(out, maxLen(A[i:], B))
	}
	for i := 0; i < n && n-i > out; i++ {
		out = Max(out, maxLen(A, B[i:]))
	}
	return out
}

func maxLen(A []int, B []int) int {
	m, n := len(A), len(B)
	out := 0
	cnt := 0
	for i := 0; i < m && i < n; i++ {
		if A[i] != B[i] {
			cnt = 0
		} else {
			cnt++
			out = Max(out, cnt)
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

// Dynamic Programming
func findLength2(A []int, B []int) int {
	m, n := len(A), len(B)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	out := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = 0
			}
			out = Max(out, dp[i+1][j+1])
		}
	}
	return out
}
