package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	out := 0
	for i, j := 0, 0; i < len(g) && j < len(s); j++ {
		if g[i] <= s[j] {
			out++
			i++
		}
	}
	return out
}
