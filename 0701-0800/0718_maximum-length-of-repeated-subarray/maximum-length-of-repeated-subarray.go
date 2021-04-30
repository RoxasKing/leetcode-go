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

// Important!

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
