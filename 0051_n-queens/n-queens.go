package main

import "strings"

/*
  n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
  给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
  每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/

func solveNQueens(n int) [][]string {
	var out [][]string
	cur := make([]string, n)
	for i := range cur {
		cur[i] = strings.Repeat(".", n)
	}
	var cols = make([]bool, n)
	var subDiagonal = make([]bool, 2*n-1)
	var mainDiagonal = make([]bool, 2*n-1)
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			tmp := make([]string, n)
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || subDiagonal[row-col+n-1] || mainDiagonal[row+col] {
				continue
			}
			tmp := []byte(cur[row])
			tmp[col] = 'Q'
			cur[row] = string(tmp)
			cols[col] = true
			mainDiagonal[row+col] = true
			subDiagonal[row-col+n-1] = true

			backtrack(row + 1)

			tmp[col] = '.'
			cur[row] = string(tmp)
			cols[col] = false
			mainDiagonal[row+col] = false
			subDiagonal[row-col+n-1] = false
		}
	}
	backtrack(0)
	return out
}
