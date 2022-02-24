package main

// Difficulty:
// Medium

// Tags:
// Simulation

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	out := make([]int, n)
	for i := range out {
		out[i] = -1
		x, y := 0, i
		for ; x < m && 0 <= y && y < n; x++ {
			if grid[x][y] == 1 && (y < n-1 && grid[x][y+1] == -1) || grid[x][y] == -1 && (y > 0 && grid[x][y-1] == 1) {
				break
			}
			if grid[x][y] == 1 {
				y++
			} else {
				y--
			}
		}
		if x == m && 0 <= y && y < n {
			out[i] = y
		}
	}
	return out
}
