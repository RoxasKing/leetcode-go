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
	head, tail, mid := 0, len(nums1), len(nums1)+(len(nums2)+1-len(nums1))>>1
	for head <= tail {
		i := head + (tail-head)>>1
		j := mid - i
		if i < tail && nums1[i] < nums2[j-1] {
			head = i + 1
		} else if i > head && nums1[i-1] > nums2[j] {
			tail = i - 1
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = Max(nums1[i-1], nums2[j-1])
			}

			if (len(nums1)+len(nums2))%2 == 1 {
				return float64(maxLeft)
			}

			var minRight int
			if i == len(nums1) {
				minRight = nums2[j]
			} else if j == len(nums2) {
				minRight = nums1[i]
			} else {
				minRight = Min(nums1[i], nums2[j])
			}

			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}
