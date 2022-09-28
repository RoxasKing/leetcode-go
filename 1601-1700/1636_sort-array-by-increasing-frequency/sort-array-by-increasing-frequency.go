package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Hash
// Sorting

func frequencySort(nums []int) []int {
	f := map[int]int{}
	a := []int{}
	for _, x := range nums {
		if f[x]++; f[x] == 1 {
			a = append(a, x)
		}
	}
	sort.Slice(a, func(i, j int) bool { return f[a[i]] < f[a[j]] || f[a[i]] == f[a[j]] && a[i] > a[j] })
	o := make([]int, 0, len(nums))
	for _, x := range a {
		for ; f[x] > 0; f[x]-- {
			o = append(o, x)
		}
	}
	return o
}
