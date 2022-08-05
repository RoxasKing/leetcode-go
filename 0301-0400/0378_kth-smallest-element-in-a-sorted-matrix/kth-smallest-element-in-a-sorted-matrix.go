package main

// Difficulty:
// Medium

// Tags:
// Binary Search

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		m := (l + r) >> 1
		c := 0
		for i, j := n-1, 0; i >= 0 && j < n; {
			if matrix[i][j] > m {
				i--
			} else {
				c += i + 1
				j++
			}
		}
		if c < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
