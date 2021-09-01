package main

// Tags:
// Binary Search

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	mid := (m + n + 1) >> 1
	i, j := 0, 0
	for l, r := 0, m; l <= r; {
		i = (l + r) >> 1
		j = mid - i
		if i < r && nums1[i] < nums2[j-1] {
			l = i + 1
		} else if i > l && nums1[i-1] > nums2[j] {
			r = i - 1
		} else {
			break
		}
	}

	maxL, minR := -1<<31, 1<<31-1
	if i > 0 && maxL < nums1[i-1] {
		maxL = nums1[i-1]
	}
	if j > 0 && maxL < nums2[j-1] {
		maxL = nums2[j-1]
	}
	if i < m && minR > nums1[i] {
		minR = nums1[i]
	}
	if j < n && minR > nums2[j] {
		minR = nums2[j]
	}

	if (m+n)&1 == 1 {
		return float64(maxL)
	}
	return float64(maxL+minR) / 2
}
