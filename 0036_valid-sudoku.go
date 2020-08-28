package main

/*
  判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
    数字 1-9 在每一行只能出现一次。
    数字 1-9 在每一列只能出现一次。
    数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
  数独部分空格内已填入了数字，空白格用 '.' 表示。
*/

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	cube := make([]map[byte]bool, 9)
	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		cube[i] = make(map[byte]bool)
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			boxIndex := (i/3)*3 + j/3
			if cube[boxIndex][board[i][j]] ||
				rows[i][board[i][j]] ||
				cols[j][board[i][j]] {
				return false
			}
			cube[boxIndex][board[i][j]] = true
			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
		}
	}
	return true
}

func isValidSudoku2(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	boxs := [9][9]bool{}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == '.' {
				continue
			}
			k := (i/3)*3 + j/3
			index := int(board[i][j] - '1')
			if rows[i][index] || cols[j][index] || boxs[k][index] {
				return false
			}
			rows[i][index] = true
			cols[j][index] = true
			boxs[k][index] = true
		}
	}
	return true
}
