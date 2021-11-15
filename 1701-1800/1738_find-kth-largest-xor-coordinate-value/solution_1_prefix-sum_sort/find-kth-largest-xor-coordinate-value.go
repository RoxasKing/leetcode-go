package main

import "sort"

// Bit Manipulation
// Prefix Sum
// Sorting

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	cnt := map[int]int{}
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = matrix[i][j]
			if i > 0 {
				f[i][j] ^= f[i-1][j]
			}
			if j > 0 {
				f[i][j] ^= f[i][j-1]
			}
			if i > 0 && j > 0 {
				f[i][j] ^= f[i-1][j-1]
			}
			cnt[f[i][j]]++
		}
	}

	arr := make([]int, 0, len(cnt))
	for num := range cnt {
		arr = append(arr, num)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	var out int
	for k > 0 {
		out = arr[0]
		k -= cnt[arr[0]]
		arr = arr[1:]
	}
	return out
}
