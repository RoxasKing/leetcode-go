package main

// Tags:
// Backtracking

func exist(board [][]byte, word string) bool {
	var dfs func(int, int, int) bool
	dfs = func(i, x, y int) bool {
		if len(word) == i {
			return true
		}
		if x < 0 || len(board)-1 < x || y < 0 || len(board[0])-1 < y || board[x][y] != word[i] {
			return false
		}
		board[x][y] = '#'
		defer func() { board[x][y] = word[i] }()
		return dfs(i+1, x-1, y) || dfs(i+1, x+1, y) || dfs(i+1, x, y-1) || dfs(i+1, x, y+1)
	}
	for i := range board {
		for j := range board[i] {
			if dfs(0, i, j) {
				return true
			}
		}
	}
	return false
}
