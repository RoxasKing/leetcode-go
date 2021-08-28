package main

// Tags:
// Dynamic Programming
// Backtracking

func escapeMaze(maze [][]string) bool {
	G = maze
	T, M, N = len(maze), len(maze[0]), len(maze[0][0])
	V = make([][][][6]bool, T)
	for i := 0; i < T; i++ {
		V[i] = make([][][6]bool, M)
		for j := 0; j < M; j++ {
			V[i][j] = make([][6]bool, N)
		}
	}
	return dp(0, 0, 0, 0)
}

var G [][]string
var T, M, N int
var V [][][][6]bool /* [t][r][c][s] */
/* s: 0-未使用永久消除术 1-已使用永久消除术并驻留 2-已使用永久消除术并跳出 | 0-未使用临时消除术， 1-已使用临时消除术 */

var dr = [4]int{-1, 1, 0, 0}
var dc = [4]int{0, 0, -1, 1}

func dp(t, r, c, s int) bool {
	if r < 0 || M-1 < r || c < 0 || N-1 < c || T <= t+M-1-r+N-1-c || V[t][r][c][s] {
		return false
	}

	if r == M-1 && c == N-1 {
		return true
	}

	V[t][r][c][s] = true

	/* 已使用永久消除术并驻留 */
	if s>>1 == 1 {
		for i := 0; i < 4; i++ {
			if dp(t+1, r+dr[i], c+dc[i], s^0b110) { /* 跳出 */
				return true
			}
		}
		return dp(t+1, r, c, s) /* 驻留 */
	}

	/* 尝试使用永久消除术 */
	if s>>1 == 0 && G[t][r][c] == '#' && dp(t, r, c, s|0b010) {
		return true
	}

	/* 使用永久消除术无法抵达终点，尝试使用临时消除术 */
	if G[t][r][c] == '#' {
		if s&1 == 1 {
			return false
		}
		s |= 1
	}

	for i := 0; i < 4; i++ {
		if dp(t+1, r+dr[i], c+dc[i], s) {
			return true
		}
	}
	return dp(t+1, r, c, s)
}
