package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func findLength(nums1 []int, nums2 []int) int {
	find := func(nums1, nums2 []int) int {
		o, m, n := 0, len(nums1), len(nums2)
		for i := 0; i < m && m-i > o; i++ {
			for j, c := 0, 0; j < n && i+j < m; j++ {
				if nums1[i+j] == nums2[j] {
					c++
					o = max(o, c)
				} else {
					c = 0
				}
			}
		}
		return o
	}
	return max(find(nums1, nums2), find(nums2, nums1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
