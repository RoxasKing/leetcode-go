package main

/*
  编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
    每行中的整数从左到右按升序排列。
    每行的第一个整数大于前一行的最后一个整数。
*/

// Binary Search
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	L, R := 0, len(matrix)*len(matrix[0])-1
	for L <= R {
		M := L + (R-L)>>1
		row := M / len(matrix[0])
		col := M % len(matrix[0])
		if matrix[row][col] < target {
			L = M + 1
		} else if matrix[row][col] > target {
			R = M - 1
		} else {
			return true
		}
	}
	return false
}
