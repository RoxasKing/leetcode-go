package main

import "sort"

// Binary Search
func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(drinks)
	out := 0
	for _, s := range staple {
		idx := sort.Search(len(drinks), func(i int) bool {
			return drinks[i] > x-s
		})
		out += idx
		out %= 1e9 + 7
	}
	return out
}
