package main

// Dfficulty:
// Easy

func projectionArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	hs := make([]int, n)
	o := 0
	for i := 0; i < m; i++ {
		h := 0
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			o++
			h = Max(h, grid[i][j])
			hs[j] = Max(hs[j], grid[i][j])
		}
		o += h
	}
	for _, h := range hs {
		o += h
	}
	return o
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
