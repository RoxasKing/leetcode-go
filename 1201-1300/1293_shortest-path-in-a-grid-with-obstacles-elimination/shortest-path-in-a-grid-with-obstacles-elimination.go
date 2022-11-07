package main

// Difficulty:
// Hard

// Tags:
// BFS
// Hash

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if m*n <= 2 {
		return m*n - 1
	}
	vis := make([][][]bool, m)
	for i := range vis {
		vis[i] = make([][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([]bool, k+1)
		}
	}
	dirs := []int{-1, 0, 1, 0, -1}
	type pair struct{ x, y, k, c int }
	for q := []pair{{0, 0, k, 0}}; len(q) > 0; q = q[1:] {
		for p, i := q[0], 0; i < 4; i++ {
			x, y, c := p.x+dirs[i], p.y+dirs[i+1], p.c+1
			if x == m-1 && y == n-1 {
				return c
			}
			if x < 0 || m-1 < x || y < 0 || n-1 < y || p.k-grid[x][y] < 0 || vis[x][y][p.k-grid[x][y]] {
				continue
			}
			vis[x][y][p.k-grid[x][y]] = true
			q = append(q, pair{x, y, p.k - grid[x][y], c})
		}
	}
	return -1
}
