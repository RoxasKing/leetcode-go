package _001_0025

import (
	"fmt"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	iMin, iMax := 0, len(nums1)
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (len(nums1)+len(nums2)+1)/2 - i
		if i < iMax && nums1[i] < nums2[j-1] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			var maxLeft float64
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = max(float64(nums1[i-1]), float64(nums2[j-1]))
			}
			if (len(nums1)+len(nums2))%2 == 1 {
				return maxLeft
			}
			var minRight float64
			if i == len(nums1) {
				minRight = float64(nums2[j])
			} else if j == len(nums2) {
				minRight = float64(nums1[i])
			} else {
				minRight = min(float64(nums1[i]), float64(nums2[j]))
			}
			return (maxLeft + minRight) / 2.0
		}
	}
	return 0.0
}

func min(i, j float64) float64 {
	if i <= j {
		return i
	}
	return j
}

func max(i, j float64) float64 {
	if i >= j {
		return i
	}
	return j
}

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
