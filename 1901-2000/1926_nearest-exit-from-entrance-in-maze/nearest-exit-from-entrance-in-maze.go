package main

// Tags:
// BFS

var fs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	x0, y0 := entrance[0], entrance[1]
	maze[x0][y0] = '+'
	q := [][3]int{{x0, y0, 0}}

	for ; len(q) > 0; q = q[1:] {
		e := q[0]
		x, y, step := e[0], e[1], e[2]

		for _, f := range fs {
			nx, ny := x+f[0], y+f[1]
			if 0 <= nx && nx < m && 0 <= ny && ny < n && maze[nx][ny] == '.' {
				if nx == 0 || nx == m-1 || ny == 0 || ny == n-1 {
					return step + 1
				}
				maze[nx][ny] = '+'
				q = append(q, [3]int{nx, ny, step + 1})
			}
		}
	}

	return -1
}
