package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting
// Monotonic Stack

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		a, b := properties[i], properties[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	o, a := 0, []int{}
	for _, p := range properties {
		for ; len(a) > 0 && a[len(a)-1] < p[1]; a = a[:len(a)-1] {
			o++
		}
		a = append(a, p[1])
	}
	return o
}
