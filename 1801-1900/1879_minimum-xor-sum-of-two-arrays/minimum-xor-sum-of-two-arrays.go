package main

import "math/bits"

// Dynamic Programming
// Bit Manipulation

func minimumXORSum(nums1 []int, nums2 []int) int {
	n := len(nums1)
	f := make([]int, 1<<n)
	for i := range f {
		f[i] = int(1e9)
	}
	f[0] = 0

	for mask := 1; mask < 1<<n; mask++ {
		c := bits.OnesCount(uint(mask))
		for i := 0; i < n; i++ {
			if mask&(1<<i) > 0 {
				f[mask] = Min(f[mask], f[mask^(1<<i)]+(nums1[c-1]^nums2[i]))
			}
		}
	}

	return f[(1<<n)-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
