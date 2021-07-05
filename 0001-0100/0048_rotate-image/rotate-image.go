package main

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n-1-i; i++ {
		for j := i; j < n-1-i; j++ {
			a, b, c, d := matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i]
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = d, a, b, c
		}
	}
}
