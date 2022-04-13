package main

// Difficulty:
// Medium

// Tags:
// Simulation

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := 0
			for x := Max(0, i-1); x <= Min(m-1, i+1); x++ {
				for y := Max(0, j-1); y <= Min(n-1, j+1); y++ {
					if x == i && y == j {
						continue
					}
					if board[x][y] == 1 || board[x][y] == 2 {
						cnt++
					}
				}
			}
			if board[i][j] == 1 && (cnt < 2 || 3 < cnt) {
				board[i][j] = 2
			}
			if board[i][j] == 0 && cnt == 3 {
				board[i][j] = 3
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
