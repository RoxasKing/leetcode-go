package main

// Tags:
// Backtracking
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, m, n, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, m, n int, word string, i, x, y int) bool {
	if i == len(word) {
		return true
	}

	if x < 0 || x > m-1 || y < 0 || y > n-1 || board[x][y] != word[i] {
		return false
	}

	ch := board[x][y]
	board[x][y] = '$'

	if dfs(board, m, n, word, i+1, x-1, y) ||
		dfs(board, m, n, word, i+1, x+1, y) ||
		dfs(board, m, n, word, i+1, x, y-1) ||
		dfs(board, m, n, word, i+1, x, y+1) {
		return true
	}

	board[x][y] = ch

	return false
}
