package crackingthecodingintervew

/*
  给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
  不占用额外内存空间能否做到？
*/

func rotate(matrix [][]int) {
	for l, r := 0, len(matrix)-1; l < r; l, r = l+1, r-1 {
		for i := l; i < r; i++ {
			matrix[l][i],
				matrix[i][r],
				matrix[len(matrix)-1-l][len(matrix)-1-i],
				matrix[len(matrix)-1-i][l] =
				matrix[len(matrix)-1-i][l],
				matrix[l][i],
				matrix[i][r],
				matrix[len(matrix)-1-l][len(matrix)-1-i]
		}
	}
}
