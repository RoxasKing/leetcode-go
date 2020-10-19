package main

/*
  给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

  示例:
    输入: 3
    输出:
    [
     [ 1, 2, 3 ],
     [ 8, 9, 4 ],
     [ 7, 6, 5 ]
    ]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/spiral-matrix-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	L, R, U, D := 0, n-1, 0, n-1
	move, index := 0, 1
	for L <= R && U <= D {
		if move == 0 {
			for i := L; i <= R; i++ {
				matrix[U][i] = index
				index++
			}
			U++
		} else if move == 1 {
			for i := U; i <= D; i++ {
				matrix[i][R] = index
				index++
			}
			R--
		} else if move == 2 {
			for i := R; i >= L; i-- {
				matrix[D][i] = index
				index++
			}
			D--
		} else {
			for i := D; i >= U; i-- {
				matrix[i][L] = index
				index++
			}
			L++
		}
		move = (move + 1) % 4
	}
	return matrix
}
