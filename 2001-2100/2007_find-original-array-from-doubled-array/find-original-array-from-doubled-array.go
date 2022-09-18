package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Greedy
// Counting
// Sorting

func findOriginalArray(changed []int) []int {
	if len(changed)&1 == 1 {
		return nil
	}
	f, a := [1e5 + 1]int{}, []int{}
	for _, x := range changed {
		if f[x]++; f[x] == 1 {
			a = append(a, x)
		}
	}
	sort.Ints(a)
	o := []int{}
	if a[0] == 0 {
		a = a[1:]
		if f[0]&1 == 1 {
			return nil
		}
		for i := 0; i < f[0]/2; i++ {
			o = append(o, 0)
		}
	}
	for _, x := range a {
		if f[x] == 0 {
			continue
		}
		if x > 50000 || f[x*2] < f[x] {
			return nil
		}
		for i := 0; i < f[x]; i++ {
			o = append(o, x)
		}
		f[x*2] -= f[x]
	}
	return o
}
