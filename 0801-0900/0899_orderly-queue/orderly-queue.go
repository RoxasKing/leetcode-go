package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Math
// Sorting

func orderlyQueue(s string, k int) string {
	n := len(s)
	if k > 1 {
		o := []byte(s)
		sort.Slice(o, func(i, j int) bool { return o[i] < o[j] })
		return string(o)
	}
	p := 0
	for i := 1; i < n; i++ {
		j := 0
		for ; j < n && s[(p+j)%n] == s[(i+j)%n]; j++ {
		}
		if s[(p+j)%n] > s[(i+j)%n] {
			p = i
		}
	}
	return s[p:] + s[:p]
}
