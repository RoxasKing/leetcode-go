package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Sorting

func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	type entry struct{ x, i int }
	entries := make([]entry, m)
	for i := 0; i < m; i++ {
		cnt := 0
		for j := 0; j < n && mat[i][j] == 1; j++ {
			cnt++
		}
		entries[i] = entry{cnt, i}
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].x < entries[j].x || entries[i].x == entries[j].x && entries[i].i < entries[j].i
	})
	out := make([]int, k)
	for i := range out {
		out[i] = entries[i].i
	}
	return out
}
