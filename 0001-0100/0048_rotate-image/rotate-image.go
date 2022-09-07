package main

// Difficulty:
// Medium

func rotate(matrix [][]int) {
	n := len(matrix)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		for i := l; i < r; i++ {
			matrix[l][i], matrix[i][r], matrix[r][n-1-i], matrix[n-1-i][l] =
				matrix[n-1-i][l], matrix[l][i], matrix[i][r], matrix[r][n-1-i]
		}
	}
}
