package main

// Difficulty:
// Hard

// Tags:
// Backtracking

func uniquePathsIII(grid [][]int) int {
	x, y, n := -1, -1, -1
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == -1 {
				continue
			}
			n++
			if grid[i][j] == 1 {
				x, y = i, j
			}
		}
	}
	return dfs(grid, x, y, n)
}

func dfs(g [][]int, x, y, n int) int {
	if x < 0 || x >= len(g) || y < 0 || y >= len(g[0]) || g[x][y] == -1 {
		return 0
	}
	if g[x][y] == 2 {
		if n == 0 {
			return 1
		}
		return 0
	}
	n--
	g[x][y] = -1
	defer func() { g[x][y] = 0 }()
	return dfs(g, x+1, y, n) + dfs(g, x-1, y, n) + dfs(g, x, y+1, n) + dfs(g, x, y-1, n)
}
