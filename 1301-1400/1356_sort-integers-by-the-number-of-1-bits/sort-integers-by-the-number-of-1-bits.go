package main

import "sort"

func sortByBits(arr []int) []int {
	n := len(arr)
	idxs := make([]int, n)
	bits := make([]int, n)
	for i := range arr {
		idxs[i] = i
		bits[i] = countBit(arr[i])
	}
	sort.Slice(idxs, func(i, j int) bool {
		if bits[idxs[i]] != bits[idxs[j]] {
			return bits[idxs[i]] < bits[idxs[j]]
		}
		return arr[idxs[i]] < arr[idxs[j]]
	})
	out := make([]int, n)
	for i := range idxs {
		out[i] = arr[idxs[i]]
	}
	return out
}

func countBit(n int) int {
	count := 0
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}
