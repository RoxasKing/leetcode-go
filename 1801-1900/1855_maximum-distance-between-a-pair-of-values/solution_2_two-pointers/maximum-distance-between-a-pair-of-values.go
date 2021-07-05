package main

// Tags:
// Two Pointers
func maxDistance(nums1 []int, nums2 []int) int {
	out := 0
	m, n := len(nums1), len(nums2)
	for i, j := m-1, n-1; i >= 0; i-- {
		for j >= 0 && nums2[j] < nums1[i] {
			j--
		}
		out = Max(out, j-i)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
