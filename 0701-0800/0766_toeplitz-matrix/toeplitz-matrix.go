package main

func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 && matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
