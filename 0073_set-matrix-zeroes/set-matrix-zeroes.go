package main

/*
  给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
*/

func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i := range matrix {
		for j := range matrix[0] {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}
