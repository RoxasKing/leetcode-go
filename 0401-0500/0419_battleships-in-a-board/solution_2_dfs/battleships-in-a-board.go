package main

// Difficulty:
// Medium

// Tags:
// DFS

func countBattleships(board [][]byte) int {
	m, n := len(board), len(board[0])
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || m-1 < i || j < 0 || n-1 < j || board[i][j] != 'X' {
			return
		}
		board[i][j] = '.'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	var out int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}
			out++
			dfs(i, j)
		}
	}
	return out
}
