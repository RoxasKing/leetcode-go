package My_LeetCode_In_Go

/*
  给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
*/

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	var row, col []int
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}
		}
	}
	for _, i := range row {
		for j := range matrix[0] {
			matrix[i][j] = 0
		}
	}
	for _, j := range col {
		for i := range matrix {
			matrix[i][j] = 0
		}
	}
}
