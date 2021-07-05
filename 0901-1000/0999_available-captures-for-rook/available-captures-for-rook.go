package main

func numRookCaptures(board [][]byte) int {
	m, n := len(board), len(board[0])
	x, y := findRook(board)
	out := 0
	for i := x + 1; i < m; i++ {
		if board[i][y] == 'B' {
			break
		} else if board[i][y] == 'p' {
			out++
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		if board[i][y] == 'B' {
			break
		} else if board[i][y] == 'p' {
			out++
			break
		}
	}
	for j := y + 1; j < n; j++ {
		if board[x][j] == 'B' {
			break
		} else if board[x][j] == 'p' {
			out++
			break
		}
	}
	for j := y - 1; j >= 0; j-- {
		if board[x][j] == 'B' {
			break
		} else if board[x][j] == 'p' {
			out++
			break
		}
	}
	return out
}

func findRook(board [][]byte) (int, int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'R' {
				return i, j
			}
		}
	}
	return -1, -1
}
