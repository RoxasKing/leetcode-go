package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting

func findClosestElements(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := abs(arr[i]-x), abs(arr[j]-x)
		return a < b || a == b && arr[i] < arr[j]
	})
	sort.Ints(arr[:k])
	return arr[:k]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
