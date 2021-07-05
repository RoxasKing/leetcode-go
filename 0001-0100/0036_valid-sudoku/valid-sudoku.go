package main

func isValidSudoku(board [][]byte) bool {
	row := [9][9]bool{}
	col := [9][9]bool{}
	box := [9][9]bool{}
	for i := range board {
		for j := range board[0] {
			if board[i][j] != '.' {
				k := (i/3)*3 + j/3
				num := board[i][j] - '1'
				if row[i][num] || col[j][num] || box[k][num] {
					return false
				}
				row[i][num] = true
				col[j][num] = true
				box[k][num] = true
			}
		}
	}
	return true
}
