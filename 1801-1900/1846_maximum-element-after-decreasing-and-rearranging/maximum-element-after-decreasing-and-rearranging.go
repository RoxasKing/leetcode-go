package main

import "sort"

// Greedy

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = Min(arr[i], arr[i-1]+1)
	}
	return arr[n-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
