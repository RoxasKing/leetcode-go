package main

// Difficulty:
// Medium

// Tags:
// Hash

func isValidSudoku(board [][]byte) bool {
	row := [9][9]bool{}
	col := [9][9]bool{}
	box := [9][9]bool{}
	for i, vs := range board {
		for j, v := range vs {
			if v != '.' {
				k := (i/3)*3 + j/3
				x := v - '1'
				if row[i][x] || col[j][x] || box[k][x] {
					return false
				}
				row[i][x] = true
				col[j][x] = true
				box[k][x] = true
			}
		}
	}
	return true
}
