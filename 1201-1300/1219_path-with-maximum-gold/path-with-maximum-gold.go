package main

// Difficulty:
// Medium

// Tags:
// Backtracking

func getMaximumGold(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	out, cur := 0, 0
	var backtrack func(int, int)
	backtrack = func(i, j int) {
		if i < 0 || m-1 < i || j < 0 || n-1 < j || grid[i][j] == 0 {
			return
		}
		gold := grid[i][j]
		grid[i][j] = 0
		cur += gold
		if out < cur {
			out = cur
		}
		backtrack(i-1, j)
		backtrack(i+1, j)
		backtrack(i, j-1)
		backtrack(i, j+1)
		cur -= gold
		grid[i][j] = gold
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backtrack(i, j)
		}
	}
	return out
}
