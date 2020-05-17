package codinginterviews

/*
  输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
  限制：
    0 <= matrix.length <= 100
    0 <= matrix[i].length <= 100
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var out []int
	var action int
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for left <= right && top <= bottom {
		switch action {
		case 0: // walk top border
			for y := left; y <= right; y++ {
				out = append(out, matrix[top][y])
			}
			top++
		case 1: // walk right border
			for x := top; x <= bottom; x++ {
				out = append(out, matrix[x][right])
			}
			right--
		case 2: // walk bottom border
			for y := right; y >= left; y-- {
				out = append(out, matrix[bottom][y])
			}
			bottom--
		case 3: // walk right border
			for x := bottom; x >= top; x-- {
				out = append(out, matrix[x][left])
			}
			left++
		}
		action = (action + 1) % 4
	}
	return out
}
