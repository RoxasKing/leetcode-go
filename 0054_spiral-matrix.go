package My_LeetCode_In_Go

/*
  给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	var state int
	out := make([]int, 0, len(matrix)*len(matrix[0]))
	for len(matrix) != 0 && len(matrix[0]) != 0 {
		switch state {
		case 0:
			for _, elem := range matrix[0] {
				out = append(out, elem)
			}
			matrix = matrix[1:]
			state = 1
		case 1:
			for i := range matrix {
				out = append(out, matrix[i][len(matrix[i])-1])
				matrix[i] = matrix[i][:len(matrix[i])-1]
			}
			state = 2
		case 2:
			for j := len(matrix[0]) - 1; j >= 0; j-- {
				out = append(out, matrix[len(matrix)-1][j])
			}
			matrix = matrix[:len(matrix)-1]
			state = 3
		case 3:
			for i := len(matrix) - 1; i >= 0; i-- {
				out = append(out, matrix[i][0])
				matrix[i] = matrix[i][1:]
			}
			state = 0
		}
	}
	return out
}
