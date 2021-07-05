package main

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	col0 := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 { // 记录第一列是否要置0
			col0 = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 // 第一列用来记录当前行是否要置0
				matrix[0][j] = 0 // 第一行用来记录当前列是否要置0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
