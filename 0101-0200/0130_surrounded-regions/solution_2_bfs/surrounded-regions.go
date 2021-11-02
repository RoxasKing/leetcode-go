package main

// Difficulty:
// Medium

// Tags:
// BFS

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	q := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
				board[i][j] = '#'
				q = append(q, []int{i, j})
			}
		}
	}
	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}
	for ; len(q) > 0; q = q[1:] {
		x, y := q[0][0], q[0][1]
		for i := 0; i < 4; i++ {
			r, c := x+dr[i], y+dc[i]
			if r < 0 || m-1 < r || c < 0 || n-1 < c || board[r][c] != 'O' {
				continue
			}
			board[r][c] = '#'
			q = append(q, []int{r, c})
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
