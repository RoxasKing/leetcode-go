package main

// Tags:
// Math
func surfaceArea(grid [][]int) int {
	n := len(grid)
	out := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			out += 2
			if i == 0 {
				out += grid[i][j]
			}
			if i == n-1 {
				out += grid[i][j]
			}
			if j == 0 {
				out += grid[i][j]
			}
			if j == n-1 {
				out += grid[i][j]
			}
			if i > 0 && grid[i][j] > grid[i-1][j] {
				out += grid[i][j] - grid[i-1][j]
			}
			if i < n-1 && grid[i][j] > grid[i+1][j] {
				out += grid[i][j] - grid[i+1][j]
			}
			if j > 0 && grid[i][j] > grid[i][j-1] {
				out += grid[i][j] - grid[i][j-1]
			}
			if j < n-1 && grid[i][j] > grid[i][j+1] {
				out += grid[i][j] - grid[i][j+1]
			}
		}
	}
	return out
}
