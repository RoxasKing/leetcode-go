package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func cherryPickup(grid [][]int) int {
	n := len(grid)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = -1 << 31
		}
	}
	f[0][0] = grid[0][0]
	for k := 1; k <= n*2-2; k++ {
		for x1 := min(k, n-1); x1 >= max(k-(n-1), 0); x1-- {
			for x2 := min(k, n-1); x2 >= x1; x2-- {
				y1, y2 := k-x1, k-x2
				if grid[x1][y1] == -1 || grid[x2][y2] == -1 {
					f[x1][x2] = -1 << 31
					continue
				}
				v := f[x1][x2] // Right, Right
				if x1 > 0 {
					v = max(v, f[x1-1][x2]) // Down, Right
				}
				if x2 > 0 {
					v = max(v, f[x1][x2-1]) // Right, Down
				}
				if x1 > 0 && x2 > 0 {
					v = max(v, f[x1-1][x2-1]) // Down, Down
				}
				v += grid[x1][y1]
				if x1 != x2 {
					v += grid[x2][y2]
				}
				f[x1][x2] = v
			}
		}
	}
	return max(f[n-1][n-1], 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
