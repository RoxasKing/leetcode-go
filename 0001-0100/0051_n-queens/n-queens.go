package main

import "strings"

// Backtracking
// DFS

func solveNQueens(n int) [][]string {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	c := make([]bool, n)
	x := make([]bool, n<<1-1)
	y := make([]bool, n<<1-1)
	out := [][]string{}
	dfs(board, c, x, y, 0, n, &out)
	return out
}

func dfs(board []string, c, x, y []bool, i, n int, out *[][]string) {
	if i == n {
		tmp := make([]string, n)
		copy(tmp, board)
		*out = append(*out, tmp)
		return
	}

	chs := []byte(board[i])
	for j := 0; j < n; j++ {
		if c[j] || x[i+j] || y[i-j+n-1] {
			continue
		}

		chs[j] = 'Q'
		board[i] = string(chs)
		c[j] = true
		x[i+j] = true
		y[i-j+n-1] = true

		dfs(board, c, x, y, i+1, n, out)

		chs[j] = '.'
		board[i] = string(chs)
		c[j] = false
		x[i+j] = false
		y[i-j+n-1] = false
	}
}
