package leetcode

/*
  n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
  给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
  每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/

func solveNQueens(n int) [][]string {
	var (
		strs         string
		out          [][]string
		track        = make([]string, n)
		cols         = make([]bool, n)
		subDiagonal  = make([]bool, 2*n-1)
		mainDiagonal = make([]bool, 2*n-1)
		backtrack    func(int)
	)
	for i := 0; i < n; i++ {
		strs += "."
	}
	for i := range track {
		track[i] = strs
	}
	place := func(row, col int) {
		tmp := []byte(track[row])
		tmp[col] = 'Q'
		track[row] = string(tmp)
		cols[col] = true
		mainDiagonal[row+col] = true
		subDiagonal[row-col+n-1] = true
	}
	unplace := func(row, col int) {
		tmp := []byte(track[row])
		tmp[col] = '.'
		track[row] = string(tmp)
		cols[col] = false
		mainDiagonal[row+col] = false
		subDiagonal[row-col+n-1] = false
	}
	addRes := func() {
		cur := make([]string, n)
		copy(cur, track)
		out = append(out, cur)
	}
	backtrack = func(row int) {
		if row == n {
			addRes()
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || subDiagonal[row-col+n-1] || mainDiagonal[row+col] {
				continue
			}
			place(row, col)
			backtrack(row + 1)
			unplace(row, col)
		}
	}
	backtrack(0)
	return out
}
