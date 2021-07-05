package main

import "sort"

// Binary Search

func kWeakestRows(mat [][]int, k int) []int {
	n, m := len(mat), len(mat[0])
	arr := make([]int, n)
	for i := range arr {
		arr[i] = sort.Search(m, func(j int) bool { return mat[i][j] == 0 })
	}

	out := make([]int, n)
	for i := range out {
		out[i] = i
	}

	sort.Slice(out, func(i, j int) bool {
		if arr[out[i]] != arr[out[j]] {
			return arr[out[i]] < arr[out[j]]
		}
		return out[i] < out[j]
	})

	return out[:k]
}
