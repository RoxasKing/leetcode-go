package leetcode

/*
  编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
    每行中的整数从左到右按升序排列。
    每行的第一个整数大于前一行的最后一个整数。
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	left, right := 0, len(matrix)*len(matrix[0])-1
	for left <= right {
		mid := left + (right-left)>>1
		row := mid / len(matrix[0])
		col := mid % len(matrix[0])
		switch {
		case target < matrix[row][col]:
			right = mid - 1
		case target > matrix[row][col]:
			left = mid + 1
		default:
			return true
		}
	}
	return false
}
