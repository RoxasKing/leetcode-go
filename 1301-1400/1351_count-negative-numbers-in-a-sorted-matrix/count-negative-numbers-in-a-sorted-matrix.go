package main

// Tags:
// Two Pointers

func countNegatives(grid [][]int) int {
	out := 0
	m, n := len(grid), len(grid[0])
	for i, j := m-1, 0; i >= 0 && j < n; {
		if grid[i][j] < 0 {
			out += n - j
			i--
		} else {
			j++
		}
	}
	return out
}
