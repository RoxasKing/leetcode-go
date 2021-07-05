package main

import "sort"

// Greedy Algorithm
// Sort

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	s := make([][]int, m)
	for i := 0; i < m; i++ {
		s[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			s[i][j] = 1
			if i > 0 && matrix[i-1][j] == 1 {
				s[i][j] += s[i-1][j]
			}
		}
	}
	out := 0
	for i := 0; i < m; i++ {
		arr := make([]int, n)
		copy(arr, s[i])
		sort.Sort(sort.Reverse(sort.IntSlice(arr)))
		for j := 0; j < n; j++ {
			out = Max(out, arr[j]*(j+1))
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
