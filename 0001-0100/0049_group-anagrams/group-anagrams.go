package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Sorting

func groupAnagrams(strs []string) [][]string {
	h := make(map[string][]string)
	for _, s := range strs {
		a := []byte(s)
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		k := string(a)
		h[k] = append(h[k], s)
	}
	o := make([][]string, 0, len(h))
	for _, v := range h {
		o = append(o, v)
	}
	return o
}
