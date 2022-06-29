package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Sorting

func minDeletions(s string) int {
	h := [1e5 + 1]int{}
	a := make([]int, 0, 26)
	for _, c := range s {
		i := int(c - 'a')
		if h[i]++; h[i] == 1 {
			a = append(a, i)
		}
	}
	sort.Slice(a, func(i, j int) bool { return h[a[i]] > h[a[j]] })
	o := 0
	for i := 1; i < len(a); i++ {
		if h[a[i-1]] <= h[a[i]] {
			val := max(0, h[a[i-1]]-1)
			o += h[a[i]] - val
			h[a[i]] = val
		}
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
