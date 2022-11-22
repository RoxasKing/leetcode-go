package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Sorting

func customSortString(order string, s string) string {
	h := [26]int{}
	vis := [26]bool{}
	idx := 26
	for _, c := range order {
		h[c-'a'] = idx
		idx--
		vis[c-'a'] = true
	}
	for c := 'a'; c <= 'z'; c++ {
		if !vis[c-'a'] {
			h[c-'a'] = idx
			idx--
		}
	}
	a := []byte(s)
	sort.Slice(a, func(i, j int) bool { return h[a[i]-'a'] > h[a[j]-'a'] })
	return string(a)
}
