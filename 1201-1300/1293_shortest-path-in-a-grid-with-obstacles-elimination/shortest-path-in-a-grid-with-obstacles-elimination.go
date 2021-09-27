package main

// Tags:
// BFS
// Queue

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
	q := []position{{0, 0, k, 0}}
	for ; len(q) > 0; q = q[1:] {
		r, c, k, steps := q[0].r, q[0].c, q[0].k, q[0].steps+1
		for i := 0; i < 4; i++ {
			x, y := r+dr[i], c+dc[i]
			if x == m-1 && y == n-1 {
				return steps
			}
			if x < 0 || m-1 < x || y < 0 || n-1 < y || k-grid[x][y] < 0 || vis[x][y][k-grid[x][y]] {
				continue
			}
			vis[x][y][k-grid[x][y]] = true
			q = append(q, position{x, y, k - grid[x][y], steps})
		}
	}
	return -1
}

type position struct{ r, c, k, steps int }

var dr = []int{1, 0, -1, 0}
var dc = []int{0, 1, 0, -1}
