package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Sorting

func arrayRankTransform(arr []int) []int {
	n := len(arr)
	h := make([]int, n)
	for i := range h {
		h[i] = i
	}
	sort.Slice(h, func(i, j int) bool { return arr[h[i]] < arr[h[j]] })
	o := make([]int, n)
	idx := 1
	for k, i := range h {
		if k > 0 && arr[h[k-1]] < arr[i] {
			idx++
		}
		o[i] = idx
	}
	return o
}
