package My_LeetCode_In_Go

/*
  n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
  上图为 8 皇后问题的一种解法。
  给定一个整数 n，返回 n 皇后不同的解决方案的数量。
*/

func totalNQueens(n int) int {
	var (
		out       int
		cols      = make([]bool, n)
		hills     = make([]bool, 2*n-1)
		dales     = make([]bool, 2*n-1)
		backtrack func(int)
	)
	backtrack = func(row int) {
		if row == n {
			out++
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || hills[row-col+n-1] || dales[row+col] {
				continue
			}
			cols[col] = true
			dales[row+col] = true
			hills[row-col+n-1] = true
			backtrack(row + 1)
			cols[col] = false
			dales[row+col] = false
			hills[row-col+n-1] = false
		}
	}
	backtrack(0)
	return out
}
