package main

// Tags:
// DFS
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, word, visited, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, visited [][]bool, i, x, y int) bool {
	if i == len(word) {
		return true
	}
	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0])-1 || visited[x][y] || board[x][y] != word[i] {
		return false
	}
	visited[x][y] = true
	out := dfs(board, word, visited, i+1, x-1, y) ||
		dfs(board, word, visited, i+1, x+1, y) ||
		dfs(board, word, visited, i+1, x, y-1) ||
		dfs(board, word, visited, i+1, x, y+1)
	visited[x][y] = false
	return out
}
