package main

// Tags:
// Binary Search
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
	for l <= r {
		mid := (l + r) / 2
		i, j := mid/n, mid%n
		if matrix[i][j] < target {
			l++
		} else if matrix[i][j] > target {
			r--
		} else {
			return true
		}
	}
	return false
}
