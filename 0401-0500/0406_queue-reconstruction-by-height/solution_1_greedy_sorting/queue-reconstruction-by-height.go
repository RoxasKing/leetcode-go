package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Greedy
// Sorting

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	o := make([][]int, len(people))
	for _, e := range people {
		copy(o[e[1]+1:], o[e[1]:])
		o[e[1]] = e
	}
	return o
}
