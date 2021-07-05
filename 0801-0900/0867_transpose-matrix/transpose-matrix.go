package main

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	out := make([][]int, n)
	for i := range out {
		out[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			out[j][i] = matrix[i][j]
		}
	}
	return out
}
