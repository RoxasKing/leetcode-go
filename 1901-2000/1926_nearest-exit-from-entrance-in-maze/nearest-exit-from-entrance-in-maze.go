package main

// DiffiCulty:
// Medium

// Tags:
// BFS

func nearestExit(maze [][]byte, entrance []int) int {
	dirs := []int{-1, 0, 1, 0, -1}
	m, n := len(maze), len(maze[0])
	x0, y0 := entrance[0], entrance[1]
	maze[x0][y0] = '+'
	for q := [][]int{{x0, y0, 0}}; len(q) > 0; q = q[1:] {
		e := q[0]
		x, y, steps := e[0], e[1], e[2]
		for i := 0; i < 4; i++ {
			nx, ny := x+dirs[i], y+dirs[i+1]
			if 0 <= nx && nx < m && 0 <= ny && ny < n && maze[nx][ny] == '.' {
				if nx == 0 || nx == m-1 || ny == 0 || ny == n-1 {
					return steps + 1
				}
				maze[nx][ny] = '+'
				q = append(q, []int{nx, ny, steps + 1})
			}
		}
	}
	return -1
}
