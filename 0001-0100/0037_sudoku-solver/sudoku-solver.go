package main

// Tags:
// Backtracking
func solveSudoku(board [][]byte) {
	row := [9][9]int{}
	col := [9][9]int{}
	box := [9][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				k := (i/3)*3 + j/3
				num := board[i][j] - '1'
				row[i][num] = 1
				col[j][num] = 1
				box[k][num] = 1
			}
		}
	}
	_ = backtrack(board, row, col, box, 0)
}

func backtrack(board [][]byte, row, col, box [9][9]int, index int) bool {
	if index == 81 {
		return true
	}
	i, j := index/9, index%9
	if board[i][j] != '.' {
		return backtrack(board, row, col, box, index+1)
	}
	k := (i/3)*3 + j/3
	for num := 0; num < 9; num++ {
		if row[i][num]+col[j][num]+box[k][num] == 0 {
			row[i][num] = 1
			col[j][num] = 1
			box[k][num] = 1

			board[i][j] = byte(num) + '1'
			if backtrack(board, row, col, box, index+1) {
				return true
			}
			board[i][j] = '.'

			row[i][num] = 0
			col[j][num] = 0
			box[k][num] = 0
		}
	}
	return false
}
