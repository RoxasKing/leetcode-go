package main

// Difficulty:
// Medium

// Tags:
// Iterator

func countBattleships(board [][]byte) int {
	out := 0
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}
			if (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				out++
			}
		}
	}
	return out
}
