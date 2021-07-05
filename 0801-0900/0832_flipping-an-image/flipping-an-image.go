package main

func flipAndInvertImage(A [][]int) [][]int {
	m, n := len(A), len(A[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n>>1; j++ {
			A[i][j], A[i][n-1-j] = A[i][n-1-j], A[i][j]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = 0
			}
		}
	}
	return A
}
