package main

import "sort"

func check(nums []int) bool {
	n := len(nums)
	arr := make([]int, n)
	copy(arr, nums)
	sort.Ints(arr)
	for i := 0; i < n; i++ {
		if isEqual(arr[i:], nums[:n-i]) && isEqual(arr[:i], nums[n-i:]) {
			return true
		}
	}
	return false
}

func isEqual(n1, n2 []int) bool {
	for i := range n1 {
		if n1[i] != n2[i] {
			return false
		}
	}
	return true
}
