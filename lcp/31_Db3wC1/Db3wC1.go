package main

// Tags:
// Dynamic Programming + DFS
func escapeMaze(maze [][]string) bool {
	t, m, n := len(maze), len(maze[0]), len(maze[0][0])
	dp := make([][][][2][2][2]bool, t) // [times][row][col][scroll1][scroll2][isStop]
	for i := range dp {
		dp[i] = make([][][2][2][2]bool, m)
		for j := range dp[i] {
			dp[i][j] = make([][2][2][2]bool, n)
		}
	}

	dfs(maze, t, m, n, dp, 0, 0, 0, 0, 0, 0)

	for i := 0; i < t; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				for l := 0; l < 2; l++ {
					if dp[i][m-1][n-1][j][k][l] {
						return true
					}
				}
			}
		}
	}
	return false
}

func dfs(maze [][]string, t, m, n int, dp [][][][2][2][2]bool, times, row, col int, scroll1, scroll2, isStop int) {
	if dp[times][row][col][scroll1][scroll2][isStop] {
		return
	}

	dp[times][row][col][scroll1][scroll2][isStop] = true

	if times == t-1 {
		return
	}

	if isStop == 1 || maze[times+1][row][col] == '.' {
		dfs(maze, t, m, n, dp, times+1, row, col, scroll1, scroll2, isStop)
	}

	for _, f := range forwards {
		i, j := row+f[0], col+f[1]
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			continue
		}
		if maze[times+1][i][j] == '.' {
			dfs(maze, t, m, n, dp, times+1, i, j, scroll1, scroll2, 0)
			continue
		}
		if scroll1 == 0 {
			dfs(maze, t, m, n, dp, times+1, i, j, 1, scroll2, 0)
		}
		if scroll2 == 0 {
			dfs(maze, t, m, n, dp, times+1, i, j, scroll1, 1, 1)
		}
	}
}

var forwards = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
