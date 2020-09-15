package main

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

// DFS + Backtracking
func solveSudoku(board [][]byte) {
	rows := make([][]int, 9)
	cols := make([][]int, 9)
	boxs := make([][]int, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make([]int, 9)
		cols[i] = make([]int, 9)
		boxs[i] = make([]int, 9)
	}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			idx := (row/3)*3 + col/3
			num := board[row][col] - '1'
			rows[row][num] = 1
			cols[col][num] = 1
			boxs[idx][num] = 1
		}
	}
	var solved bool
	var backTrack func(int)
	backTrack = func(index int) {
		if index == 81 {
			solved = true
			return
		}
		row, col := index/9, index%9
		if board[row][col] != '.' {
			backTrack(index + 1)
			return
		}
		idx := (row/3)*3 + col/3
		for i := 0; i < 9; i++ {
			if rows[row][i]+cols[col][i]+boxs[idx][i] == 0 {
				rows[row][i] = 1
				cols[col][i] = 1
				boxs[idx][i] = 1
				board[row][col] = byte(i) + '1'
				backTrack(index + 1)
				if solved {
					return
				}
				rows[row][i] = 0
				cols[col][i] = 0
				boxs[idx][i] = 0
				board[row][col] = '.'
			}
		}
	}
	backTrack(0)
}
