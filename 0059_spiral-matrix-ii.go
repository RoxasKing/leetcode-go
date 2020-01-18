package My_LeetCode_In_Go

/*
  给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
*/

func generateMatrix(n int) [][]int {
	out := make([][]int, n)
	for i := range out {
		out[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	sum := n * n
	var status int
	for i := 0; i != sum; {
		switch status {
		case 0:
			for j := left; j <= right; j++ {
				i++
				out[top][j] = i
			}
			top++
			status = 1
		case 1:
			for j := top; j <= bottom; j++ {
				i++
				out[j][right] = i
			}
			right--
			status = 2
		case 2:
			for j := right; j >= left; j-- {
				i++
				out[bottom][j] = i
			}
			bottom--
			status = 3
		case 3:
			for j := bottom; j >= top; j-- {
				i++
				out[j][left] = i
			}
			left++
			status = 0
		}
	}
	return out
}
