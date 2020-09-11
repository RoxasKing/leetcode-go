package main

/*
  给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	const (
		UP    = 0
		DOWN  = 1
		LEFT  = 2
		RIGHT = 3
	)
	rows, cols := len(matrix), len(matrix[0])
	U, D, L, R := 0, rows-1, 0, cols-1
	move := RIGHT
	out := make([]int, 0, rows*cols)
	for U <= D && L <= R {
		switch move {
		case UP:
			for i := D; i >= U; i-- {
				out = append(out, matrix[i][L])
			}
			L++
			move = RIGHT
		case DOWN:
			for i := U; i <= D; i++ {
				out = append(out, matrix[i][R])
			}
			R--
			move = LEFT
		case LEFT:
			for j := R; j >= L; j-- {
				out = append(out, matrix[D][j])
			}
			D--
			move = UP
		case RIGHT:
			for j := L; j <= R; j++ {
				out = append(out, matrix[U][j])
			}
			U++
			move = DOWN
		}
	}
	return out
}
