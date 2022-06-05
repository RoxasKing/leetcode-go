package main

// Difficulty:
// Hard

// Tags:
// Backtracking

func totalNQueens(n int) int {
	col := make([]bool, n)
	pos := make([]bool, n*2-1)
	neg := make([]bool, n*2-1)
	o := 0
	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			o++
		}
		for j := 0; j < n; j++ {
			if col[j] || pos[i-j+n-1] || neg[i+j] {
				continue
			}
			col[j] = true
			pos[i-j+n-1] = true
			neg[i+j] = true
			backtrack(i + 1)
			col[j] = false
			pos[i-j+n-1] = false
			neg[i+j] = false
		}
	}
	backtrack(0)
	return o
}
