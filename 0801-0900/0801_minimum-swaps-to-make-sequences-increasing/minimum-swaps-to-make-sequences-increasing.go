package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func minSwap(nums1 []int, nums2 []int) int {
	f0, f1 := 0, 1
	for i := 1; i < len(nums1); i++ {
		t0, t1 := 1<<31-1, 1<<31-1
		if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			t0, t1 = min(t0, f0), min(t1, f1+1)
		}
		if nums1[i-1] < nums2[i] && nums2[i-1] < nums1[i] {
			t0, t1 = min(t0, f1), min(t1, f0+1)
		}
		f0, f1 = t0, t1
	}
	return min(f0, f1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
