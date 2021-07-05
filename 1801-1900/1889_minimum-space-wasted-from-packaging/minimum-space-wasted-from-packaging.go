package main

import "sort"

// Prefix Sum
// Binary Search

func minWastedSpace(packages []int, boxes [][]int) int {
	n := len(packages)
	sort.Ints(packages)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + packages[i]
	}

	out := -1
	for _, bs := range boxes {
		sort.Ints(bs)
		k := len(bs)
		if bs[k-1] < packages[n-1] {
			continue
		}
		sum := 0
		for i, l, r := 0, 0, 0; i < k; i, l = i+1, r {
			r = sort.Search(n, func(j int) bool { return packages[j] > bs[i] })
			sum += bs[i]*(r-l) - (pSum[r] - pSum[l])
		}
		if out == -1 || sum < out {
			out = sum
		}
	}
	return out % (1e9 + 7)
}
