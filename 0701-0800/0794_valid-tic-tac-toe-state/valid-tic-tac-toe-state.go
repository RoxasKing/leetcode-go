package main

// Difficulty:
// Medium

func validTicTacToe(board []string) bool {
	cntX, cntO := 0, 0
	for i := 0; i < 9; i++ {
		if board[i/3][i%3] == 'X' {
			cntX++
		} else if board[i/3][i%3] == 'O' {
			cntO++
		}
	}
	return cntX == cntO && !win(board, 'X') || cntX == cntO+1 && !win(board, 'O')
}

func win(board []string, p byte) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == p && board[i][1] == p && board[i][2] == p ||
			board[0][i] == p && board[1][i] == p && board[2][i] == p {
			return true
		}
	}
	return board[0][0] == p && board[1][1] == p && board[2][2] == p ||
		board[0][2] == p && board[1][1] == p && board[2][0] == p
}
