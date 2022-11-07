package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Greedy
// Sorting

func earliestFullBloom(plantTime []int, growTime []int) int {
	type pair struct{ pt, gt int }
	a := make([]pair, len(plantTime))
	for i, pt := range plantTime {
		a[i] = pair{pt, growTime[i]}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].pt+max(a[i].gt, a[j].pt+a[j].gt) < a[j].pt+max(a[j].gt, a[i].pt+a[i].gt)
	})
	o, c := 0, 0
	for _, p := range a {
		c += p.pt
		o = max(o, c+p.gt)
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
