package main

// Tags:
// Binary Search
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		limit := (l + r) >> 1
		if countSmaller(matrix, n, limit) < k {
			l = limit + 1
		} else {
			r = limit
		}
	}
	return l
}

func countSmaller(matrix [][]int, n, limit int) int {
	cnt := 0
	for i, j := n-1, 0; i >= 0 && j < n; {
		if matrix[i][j] <= limit {
			cnt += i + 1
			j++
		} else {
			i--
		}
	}
	return cnt
}
