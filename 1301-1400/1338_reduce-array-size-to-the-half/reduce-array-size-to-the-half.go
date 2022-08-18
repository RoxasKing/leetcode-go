package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Greedy
// Sorting

func minSetSize(arr []int) int {
	f := [1e5 + 1]int{}
	a := []int{}
	for _, x := range arr {
		if f[x]++; f[x] == 1 {
			a = append(a, x)
		}
	}
	sort.Slice(a, func(i, j int) bool { return f[a[i]] > f[a[j]] })
	o := 0
	for c := 0; c < len(arr)/2; c, a = c+f[a[0]], a[1:] {
		o++
	}
	return o
}
