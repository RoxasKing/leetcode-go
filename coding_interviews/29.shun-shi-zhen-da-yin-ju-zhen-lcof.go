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
	var cur int
	for len(matrix) > cur<<1 && len(matrix[0]) > cur<<1 {
		endRow := len(matrix) - 1 - cur
		endCol := len(matrix[0]) - 1 - cur
		for i := cur; i <= endCol; i++ {
			out = append(out, matrix[cur][i])
		}
		if cur < endRow {
			for i := cur + 1; i <= endRow; i++ {
				out = append(out, matrix[i][endCol])
			}
		}
		if cur < endRow && cur < endCol {
			for i := endCol - 1; i >= cur; i-- {
				out = append(out, matrix[endRow][i])
			}
		}
		if cur < endRow-1 && cur < endCol {
			for i := endRow - 1; i >= cur+1; i-- {
				out = append(out, matrix[i][cur])
			}
		}
		cur++
	}
	return out
}
