package main

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	for i := 0; i < 4; i++ {
		if equal(mat, target, n) {
			return true
		}
		rotate(mat, n)
	}
	return false
}

func rotate(matrix [][]int, n int) {
	for i := 0; i < n-1-i; i++ {
		for j := i; j < n-1-i; j++ {
			a, b, c, d := matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i]
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = d, a, b, c
		}
	}
}

func equal(mat, target [][]int, n int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}
	return true
}
