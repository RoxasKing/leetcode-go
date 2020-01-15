package My_LeetCode_In_Go

/*
  给定一个 n × n 的二维矩阵表示一个图像。
  将图像顺时针旋转 90 度。
  说明：
  你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
*/

func rotate(matrix [][]int) {
	head, tail := 0, len(matrix)-1
	for head < tail {
		for i := head; i < tail; i++ {
			matrix[head][i],
				matrix[i][tail],
				matrix[len(matrix)-1-head][len(matrix)-1-i],
				matrix[len(matrix)-1-i][head] =
				matrix[len(matrix)-1-i][head],
				matrix[head][i],
				matrix[i][tail],
				matrix[len(matrix)-1-head][len(matrix)-1-i]
		}
		head, tail = head+1, tail-1
	}
}
