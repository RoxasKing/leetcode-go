package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Sorting
// Binary Search

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[1] < b[1] || a[1] == b[1] && a[0] > b[0]
	})
	arr := []int{}
	for _, e := range envelopes {
		if i := sort.Search(len(arr), func(i int) bool { return arr[i] >= e[0] }); i < len(arr) {
			arr[i] = e[0]
			continue
		}
		arr = append(arr, e[0])
	}
	return len(arr)
}
