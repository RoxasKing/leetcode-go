package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	for x, y := m-1, 0; x >= 0 && y < n; {
		if matrix[x][y] < target {
			y++
		} else if matrix[x][y] > target {
			x--
		} else {
			return true
		}
	}
	return false
}
