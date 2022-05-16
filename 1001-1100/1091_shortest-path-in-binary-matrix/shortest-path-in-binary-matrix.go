package main

// Difficulty:
// Medium

// Tags:
// BFS
// Hash

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	vis[0][0] = true
	o := 0
	for q := [][]int{{0, 0}}; len(q) > 0; {
		o++
		m := len(q)
		for i := 0; i < m; i++ {
			x0, y0 := q[i][0], q[i][1]
			if x0 == n-1 && y0 == n-1 {
				return o
			}
			for x := max(0, x0-1); x <= min(n-1, x0+1); x++ {
				for y := max(0, y0-1); y <= min(n-1, y0+1); y++ {
					if grid[x][y] == 1 || x == x0 && y == y0 || vis[x][y] {
						continue
					}
					vis[x][y] = true
					q = append(q, []int{x, y})
				}
			}
		}
		q = q[m:]
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
