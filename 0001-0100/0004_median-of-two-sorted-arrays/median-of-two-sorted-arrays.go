package main

// Tags:
// Binary Search
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	n1, n2 := len(nums1), len(nums2)

	l1, r1 := 0, n1
	m, m1, m2 := (n1+n2+1)>>1, 0, 0
	for l1 <= r1 {
		m1 = (l1 + r1) >> 1
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
