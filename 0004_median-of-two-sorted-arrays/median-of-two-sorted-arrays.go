package main

/*
  给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
  进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

  示例 1：
    输入：nums1 = [1,3], nums2 = [2]
    输出：2.00000
    解释：合并数组 = [1,2,3] ，中位数 2

  示例 2：
    输入：nums1 = [1,2], nums2 = [3,4]
    输出：2.50000
    解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
  示例 3：
    输入：nums1 = [0,0], nums2 = [0,0]
    输出：0.00000
  示例 4：
    输入：nums1 = [], nums2 = [1]
    输出：1.00000
  示例 5：
    输入：nums1 = [2], nums2 = []
    输出：2.00000

  提示：
    nums1.length == m
    nums2.length == n
    0 <= m <= 1000
    0 <= n <= 1000
    1 <= m + n <= 2000
    -106 <= nums1[i], nums2[i] <= 106

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	n1, n2 := len(nums1), len(nums2)

	l1, r1 := 0, n1
	m, m1, m2 := n1+(n2-n1+1)>>1, 0, 0
	for l1 <= r1 {
		m1 = l1 + (r1-l1)>>1
		m2 = m - m1
		if m1 < r1 && nums1[m1] < nums2[m2-1] {
			l1 = m1 + 1
		} else if m1 > l1 && nums1[m1-1] > nums2[m2] {
			r1 = m1 - 1
		} else {
			break
		}
	}

	if n1 == n2 {
		if m1 == 0 {
			return float64(nums2[n2-1]+nums1[0]) / 2
		} else if m2 == 0 {
			return float64(nums1[n1-1]+nums2[0]) / 2
		}
		return float64(Max(nums1[m1-1], nums2[m2-1])+Min(nums1[m1], nums2[m2])) / 2
	}

	maxL := nums2[m2-1]
	if m1 > 0 {
		maxL = Max(maxL, nums1[m1-1])
	}
	if (n1+n2)&1 == 1 {
		return float64(maxL)
	}
	minR := nums2[m2]
	if m1 < n1 {
		minR = Min(minR, nums1[m1])
	}
	return float64(maxL+minR) / 2
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
