package main

// Difficulty:
// Medium

// Tags:
// Simulation

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	o := make([]int, 0, m*n)
	for i, j, x := 0, 0, 1; len(o) < m*n; x = -x {
		o = append(o, mat[i][j])
		for 0 <= i-x && i-x < m && 0 <= j+x && j+x < n {
			i -= x
			j += x
			o = append(o, mat[i][j])
		}
		if 0 < i && i < m-1 || i == 0 && j == n-1 {
			i++
		} else {
			j++
		}
	}
	return o
}
