package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Sorting

func suggestedProducts(products []string, searchWord string) [][]string {
	n := len(searchWord)
	o := make([][]string, 0, n)
	for i := 1; i <= n; i++ {
		prefix := searchWord[:i]
		a := []string{}
		for _, s := range products {
			if len(s) < i || prefix != s[:i] {
				continue
			}
			if a = append(a, s); len(a) > 3 {
				sort.Strings(a)
				a = a[:3]
			}
		}
		sort.Strings(a)
		o = append(o, a)
	}
	return o
}
