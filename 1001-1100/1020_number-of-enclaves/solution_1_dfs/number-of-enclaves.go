package main

// Difficulty:
// Medium

// Tags:
// DFS

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	out, cur := 0, 0
	valid := true
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || m-1 < i || j < 0 || n-1 < j {
			valid = false
			return
		}
		if grid[i][j] == 0 {
			return
		}
		cur++
		grid[i][j] = 0
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			cur = 0
			valid = true
			dfs(i, j)
			if valid {
				out += cur
			}
		}
	}
	return out
}
