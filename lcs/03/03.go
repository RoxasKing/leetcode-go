package main

// Tags:
// DFS

func largestArea(grid []string) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				visited[i][j] = true
				if i+1 < m && grid[i+1][j] != '0' {
					walk(grid, visited, m, n, i+1, j, grid[i+1][j])
				}
				if i-1 >= 0 && grid[i-1][j] != '0' {
					walk(grid, visited, m, n, i-1, j, grid[i-1][j])
				}
				if j+1 < n && grid[i][j+1] != '0' {
					walk(grid, visited, m, n, i, j+1, grid[i][j+1])
				}
				if j-1 >= 0 && grid[i][j-1] != '0' {
					walk(grid, visited, m, n, i, j-1, grid[i][j-1])
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		if !visited[i][0] {
			walk(grid, visited, m, n, i, 0, grid[i][0])
		}
		if !visited[i][n-1] {
			walk(grid, visited, m, n, i, n-1, grid[i][n-1])
		}
	}
	for j := 0; j < n; j++ {
		if !visited[0][j] {
			walk(grid, visited, m, n, 0, j, grid[0][j])
		}
		if !visited[m-1][j] {
			walk(grid, visited, m, n, m-1, j, grid[m-1][j])
		}
	}

	out := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] {
				continue
			}
			q := [][]int{{i, j}}
			visited[i][j] = true
			cnt := 0
			for len(q) > 0 {
				x, y := q[0][0], q[0][1]
				q = q[1:]
				cnt++
				if x+1 < m && !visited[x+1][y] && grid[x+1][y] == grid[x][y] {
					visited[x+1][y] = true
					q = append(q, []int{x + 1, y})
				}
				if x-1 >= 0 && !visited[x-1][y] && grid[x-1][y] == grid[x][y] {
					visited[x-1][y] = true
					q = append(q, []int{x - 1, y})
				}
				if y+1 < n && !visited[x][y+1] && grid[x][y+1] == grid[x][y] {
					visited[x][y+1] = true
					q = append(q, []int{x, y + 1})
				}
				if y-1 >= 0 && !visited[x][y-1] && grid[x][y-1] == grid[x][y] {
					visited[x][y-1] = true
					q = append(q, []int{x, y - 1})
				}
			}
			out = Max(out, cnt)
		}
	}
	return out
}

func walk(grid []string, visited [][]bool, m, n, i, j int, val byte) {
	if i < 0 || i > m-1 || j < 0 || j > n-1 || visited[i][j] || grid[i][j] != val {
		return
	}
	visited[i][j] = true
	walk(grid, visited, m, n, i+1, j, val)
	walk(grid, visited, m, n, i, j+1, val)
	walk(grid, visited, m, n, i-1, j, val)
	walk(grid, visited, m, n, i, j-1, val)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
