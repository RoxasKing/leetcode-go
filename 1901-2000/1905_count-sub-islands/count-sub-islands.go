package main

// Tags:
// DFS

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	out := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 0 || grid1[i][j] == 0 {
				continue
			}
			grid2[i][j] = 0
			if isSubIsland(grid1, grid2, m, n, i, j) {
				out++
			}
		}
	}
	return out
}

func isSubIsland(grid1, grid2 [][]int, m, n, i, j int) bool {
	out := true
	for _, e := range moves {
		x, y := i+e[0], j+e[1]
		if x < 0 || x > m-1 || y < 0 || y > n-1 || grid2[x][y] == 0 {
			continue
		}
		grid2[x][y] = 0
		out = grid1[x][y] == 1 && out
		out = isSubIsland(grid1, grid2, m, n, x, y) && out
	}
	return out
}

var moves = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
