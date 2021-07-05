package main

// Tags:
// BFS
func orangesRotting(grid [][]int) int {
	forwards := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid), len(grid[0])
	q := [][]int{}
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				cnt++
			}
		}
	}
	out := 0
	for len(q) > 0 {
		size := len(q)
		for _, e := range q {
			x, y := e[0], e[1]
			for _, f := range forwards {
				nx, ny := x+f[0], y+f[1]
				if nx < 0 || m-1 < nx || ny < 0 || n-1 < ny || grid[nx][ny] != 1 {
					continue
				}
				grid[nx][ny] = 2
				cnt--
				q = append(q, []int{nx, ny})
			}
		}
		q = q[size:]
		if len(q) > 0 {
			out++
		}
	}
	if cnt > 0 {
		return -1
	}
	return out
}
