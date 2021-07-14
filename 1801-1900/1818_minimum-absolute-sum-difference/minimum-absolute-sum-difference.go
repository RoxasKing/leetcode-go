package main

import (
	"sort"
)

// Tags:
// Binary Search

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n := len(nums1)
	arr := make([]int, n)
	copy(arr, nums1)
	sort.Ints(arr)

	var max int
	for i := 0; i < n; i++ {
		cur := Abs(nums1[i] - nums2[i])
		j := sort.Search(n, func(j int) bool { return arr[j] >= nums2[i] })
		if j < n && cur-Abs(arr[j]-nums2[i]) > max {
			max = cur - Abs(arr[j]-nums2[i])
		}
		if j > 0 && cur-Abs(arr[j-1]-nums2[i]) > max {
			max = cur - Abs(arr[j-1]-nums2[i])
		}
	}

	var out int
	for i := 0; i < n; i++ {
		out += Abs(nums1[i] - nums2[i])
		out %= mod
	}

	return (out - max + mod) % mod
}

var mod = int(1e9 + 7)

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
