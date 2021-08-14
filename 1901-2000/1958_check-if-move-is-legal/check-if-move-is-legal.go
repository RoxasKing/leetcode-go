package main

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	m, n := len(board), len(board[0])
	return isValid(board, rMove, cMove, color, m, n, 1, 0) ||
		isValid(board, rMove, cMove, color, m, n, -1, 0) ||
		isValid(board, rMove, cMove, color, m, n, 0, 1) ||
		isValid(board, rMove, cMove, color, m, n, 0, -1) ||
		isValid(board, rMove, cMove, color, m, n, 1, 1) ||
		isValid(board, rMove, cMove, color, m, n, 1, -1) ||
		isValid(board, rMove, cMove, color, m, n, -1, 1) ||
		isValid(board, rMove, cMove, color, m, n, -1, -1)
}

func isValid(board [][]byte, x int, y int, color byte, m, n, a, b int) bool {
	cnt := 0
	for 0 <= x+a && x+a < m && 0 <= y+b && y+b < n {
		if board[x+a][y+b] == 'W' && color == 'B' || board[x+a][y+b] == 'B' && color == 'W' {
			cnt++
			x += a
			y += b
			continue
		}
		break
	}
	if cnt == 0 {
		return false
	}

	cnt = 0
	for 0 <= x+a && x+a < m && 0 <= y+b && y+b < n {
		if board[x+a][y+b] == color {
			cnt++
			x += a
			y += b
			continue
		}
		break
	}
	return cnt > 0
}
