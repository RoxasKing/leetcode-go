package My_LeetCode_In_Go

/*
  编写一个程序，通过已填充的空格来解决数独问题。
  一个数独的解法需遵循如下规则：
    数字 1-9 在每一行只能出现一次。
    数字 1-9 在每一列只能出现一次。
    数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
  空白格用 '.' 表示。
  Note:
    给定的数独序列只包含数字 1-9 和字符 '.' 。
    你可以假设给定的数独只有唯一解。
    给定数独永远是 9x9 形式的。
*/

func solveSudoku(board [][]byte) {
	rows := make([][]int, 9)
	cols := make([][]int, 9)
	boxs := make([][]int, 9)
	for i := range rows {
		rows[i] = make([]int, 10)
		cols[i] = make([]int, 10)
		boxs[i] = make([]int, 10)
	}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			idx := (row/3)*3 + col/3
			num := board[row][col] - '0'
			rows[row][num] = 1
			cols[col][num] = 1
			boxs[idx][num] = 1
		}
	}
	var (
		solved bool
		next   func(row, col int)
		solve  func(row, col int)
	)
	next = func(row, col int) {
		if col == 8 && row == 8 {
			solved = true
		} else {
			if col == 8 {
				solve(row+1, 0)
			} else {
				solve(row, col+1)
			}
		}
	}
	solve = func(row, col int) {
		if board[row][col] == '.' {
			for i := 1; i <= 9; i++ {
				idx := (row/3)*3 + col/3
				if rows[row][i]+cols[col][i]+boxs[idx][i] == 0 {
					rows[row][i] = 1
					cols[col][i] = 1
					boxs[idx][i] = 1
					board[row][col] = byte(i + '0')
					next(row, col)
					if !solved {
						rows[row][i] = 0
						cols[col][i] = 0
						boxs[idx][i] = 0
						board[row][col] = '.'
					}
				}
			}
		} else {
			next(row, col)
		}
	}
	solve(0, 0)
}
