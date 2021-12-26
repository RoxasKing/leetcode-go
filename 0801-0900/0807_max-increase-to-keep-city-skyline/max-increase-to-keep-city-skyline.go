package main

// Difficulty:
// Medium

// Tags:
// Greedy

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rMax, cMax := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rMax[i] = Max(rMax[i], grid[i][j])
			cMax[j] = Max(cMax[j], grid[i][j])
		}
	}
	var out int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			out += Min(rMax[i], cMax[j]) - grid[i][j]
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
