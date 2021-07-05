package main

func islandPerimeter(grid [][]int) int {
	var out int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				continue
			}
			for _, step := range steps {
				newI, newJ := i+step[0], j+step[1]
				if newI < 0 || len(grid)-1 < newI || newJ < 0 || len(grid[0])-1 < newJ || grid[newI][newJ] == 0 {
					out++
				}
			}
		}
	}
	return out
}

// DFS
func islandPerimeter2(grid [][]int) int {
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				var out int
				dfs(grid, i, j, &out)
				return out
			}
		}
	}
	return 0
}

func dfs(grid [][]int, row, col int, out *int) {
	grid[row][col] = -1
	for _, step := range steps {
		newR, newC := row+step[0], col+step[1]
		if newR < 0 || len(grid)-1 < newR || newC < 0 || len(grid[0])-1 < newC || grid[newR][newC] == 0 {
			*out++
			continue
		}
		if grid[newR][newC] == 1 {
			dfs(grid, newR, newC, out)
		}
	}
}

var steps = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
