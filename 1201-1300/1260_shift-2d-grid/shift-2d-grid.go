package main

// Difficulty:
// Easy

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	o := make([][]int, m)
	for i := range o {
		o[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x, y := (i+(j+k)/n)%m, (j+k)%n
			o[x][y] = grid[i][j]
		}
	}
	return o
}
