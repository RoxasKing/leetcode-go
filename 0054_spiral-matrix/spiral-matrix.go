package main

/*
  给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

  示例 1:
    输入:
    [
     [ 1, 2, 3 ],
     [ 4, 5, 6 ],
     [ 7, 8, 9 ]
    ]
    输出: [1,2,3,6,9,8,7,4,5]

  示例 2:
    输入:
    [
      [1, 2, 3, 4],
      [5, 6, 7, 8],
      [9,10,11,12]
    ]
    输出: [1,2,3,4,8,12,11,10,9,5,6,7]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/spiral-matrix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	L, R, U, D := 0, n-1, 0, m-1
	out := make([]int, 0, m*n)
	move := 0 // 0: left->right, 1: up->down, 2: right->left, 3: down to up
	for L <= R && U <= D {
		if move == 0 {
			out = append(out, matrix[U][L:R+1]...)
			U++
		} else if move == 1 {
			for i := U; i <= D; i++ {
				out = append(out, matrix[i][R])
			}
			R--
		} else if move == 2 {
			for j := R; j >= L; j-- {
				out = append(out, matrix[D][j])
			}
			D--
		} else {
			for i := D; i >= U; i-- {
				out = append(out, matrix[i][L])
			}
			L++
		}
		move = (move + 1) % 4
	}
	return out
}
