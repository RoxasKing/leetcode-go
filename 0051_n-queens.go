package leetcode

import "strings"

/*
  n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
  给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
  每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/

func solveNQueens(n int) [][]string {
	var (
		out       [][]string
		track     = make([]string, n)
		cols      = make([]bool, n)
		hills     = make([]bool, 2*n-1)
		dales     = make([]bool, 2*n-1)
		backtrack func(int)
	)
	for i := range track {
		track[i] = strings.Repeat(".", n)
	}
	backtrack = func(row int) {
		if row == n {
			out = append(out, append(make([]string, 0, n), track...))
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || hills[row-col+n-1] || dales[row+col] {
				continue
			}

			tmp := []byte(track[row])
			tmp[col] = 'Q'
			track[row] = string(tmp)

			cols[col] = true
			dales[row+col] = true
			hills[row-col+n-1] = true

			backtrack(row + 1)

			cols[col] = false
			dales[row+col] = false
			hills[row-col+n-1] = false

			tmp[col] = '.'
			track[row] = string(tmp)
		}
	}
	backtrack(0)
	return out
}
