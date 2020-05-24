package leetcode

/*
  给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
  请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
  你可以假设 nums1 和 nums2 不会同时为空。
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	l, r, m := 0, len(nums1), (len(nums2)+len(nums1)+1)>>1
	var m1, m2 int
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

	var maxL int
	if m1 == 0 {
		maxL = nums2[m2-1]
	} else if m2 == 0 {
		maxL = nums1[m1-1]
	} else {
		maxL = Max(nums1[m1-1], nums2[m2-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(maxL)
	}

	var minR int
	if m1 == len(nums1) {
		minR = nums2[m2]
	} else if m2 == len(nums2) {
		minR = nums1[m1]
	} else {
		minR = Min(nums1[m1], nums2[m2])
	}
	return float64(maxL+minR) / 2
}
