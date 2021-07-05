package main

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	locates := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i := range board {
		for j := range board[0] {
			var count int
			for _, locate := range locates {
				newI, newJ := i+locate[0], j+locate[1]
				if 0 <= newI && newI < len(board) &&
					0 <= newJ && newJ < len(board[0]) &&
					(board[newI][newJ] == 1 || board[newI][newJ] == 3) {
					count++
				}
			}
			if board[i][j] == 1 && (count < 2 || count > 3) {
				board[i][j] = 3
			} else if board[i][j] == 0 && count == 3 {
				board[i][j] = 4
			}
		}
	}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 3 {
				board[i][j] = 0
			} else if board[i][j] == 4 {
				board[i][j] = 1
			}
		}
	}
}
