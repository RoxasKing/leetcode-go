package main

/*
  给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
  请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
  你可以假设 nums1 和 nums2 不会同时为空。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) { // make len(nums1) < len(nums2)
		nums1, nums2 = nums2, nums1
	}

	var (
		l = 0
		r = len(nums1)
		m = (len(nums1) + len(nums2) + 1) >> 1
		// if (len(nums1)+len(nums2)) % 2 = 1
		// makes nums1 and nums2's left half bigger than right half
		m1 int // nums1's right half's first num
		m2 int // nums2's right half's first num
	)

	for l <= r {
		m1 = (l + r) >> 1
		m2 = m - m1
		if m1 < r && nums1[m1] < nums2[m2-1] {
			l = m1 + 1
		} else if m1 > l && nums1[m1-1] > nums2[m2] {
			r = m1 - 1
		} else {
			break
		}
	}

	if len(nums1) == len(nums2) {
		if m1 == 0 { // all nums1's num is bigger than nums2's num
			return float64(nums1[0]+nums2[len(nums2)-1]) / 2
		} else if m2 == 0 { // all nums2's num is bigger than nums1's num
			return float64(nums1[len(nums1)-1]+nums2[0]) / 2
		}
		return float64(Max(nums1[m1-1], nums2[m2-1])+Min(nums1[m1], nums2[m2])) / 2
	}

	maxL := nums2[m2-1]
	if m1 > 0 {
		maxL = Max(maxL, nums1[m1-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(maxL)
	}
	minR := nums2[m2]
	if m1 < len(nums1) {
		minR = Min(minR, nums1[m1])
	}
	return float64(maxL+minR) / 2
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
