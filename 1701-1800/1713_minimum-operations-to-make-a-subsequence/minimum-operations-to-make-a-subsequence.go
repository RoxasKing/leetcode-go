package main

import (
	"sort"
)

// Tags:
// Binary Search
// Hash
// Greedy Algoritm

func minOperations(target []int, arr []int) int {
	dict := map[int]int{}
	for i := range target {
		dict[target[i]] = i
	}
	idxs := []int{}
	for i := range arr {
		if j, ok := dict[arr[i]]; ok {
			k := sort.SearchInts(idxs, j)
			if k < len(idxs) {
				idxs[k] = j
			} else {
				idxs = append(idxs, j)
			}
		}
	}
	return len(target) - len(idxs)
}
