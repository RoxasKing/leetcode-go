package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting

func groupThePeople(groupSizes []int) [][]int {
	type pair struct{ i, g int }
	a := []pair{}
	for i, g := range groupSizes {
		a = append(a, pair{i, g})
	}
	sort.SliceStable(a, func(i, j int) bool { return a[i].g < a[j].g })
	o := [][]int{}
	for len(a) > 0 {
		c := a[0].g
		t := []int{}
		for ; c > 0; c-- {
			t = append(t, a[0].i)
			a = a[1:]
		}
		o = append(o, t)
	}
	return o
}
