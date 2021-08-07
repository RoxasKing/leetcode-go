package main

// Tags:
// Backtracking

func exist(board [][]byte, word string) bool {
	m, n, k := len(board), len(board[0]), len(word)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(board, word, m, n, k, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, m, n, k int, i, x, y int) bool {
	if i == k {
		return true
	}

	if x < 0 || x > m-1 || y < 0 || y > n-1 || board[x][y] != word[i] {
		return false
	}

	ch := board[x][y]
	defer func() { board[x][y] = ch }()
	board[x][y] = ' '

	if backtrack(board, word, m, n, k, i+1, x-1, y) ||
		backtrack(board, word, m, n, k, i+1, x+1, y) ||
		backtrack(board, word, m, n, k, i+1, x, y-1) ||
		backtrack(board, word, m, n, k, i+1, x, y+1) {
		return true
	}
	return false
}
