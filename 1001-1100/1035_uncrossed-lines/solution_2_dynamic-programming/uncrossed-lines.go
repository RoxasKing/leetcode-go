package main

// Tags:
// Dynamic Programming

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	f0 := make([]int, n2+1)
	f1 := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			f1[j+1] = f0[j]
			if nums1[i] == nums2[j] {
				f1[j+1]++
			}
			f1[j+1] = Max(f1[j+1], Max(f0[j+1], f1[j]))
		}
		f0, f1 = f1, f0
	}
	return f0[n2]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
