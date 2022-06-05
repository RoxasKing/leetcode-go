package main

import "strings"

// Difficulty:
// Hard

// Tags:
// Backtracking

func solveNQueens(n int) [][]string {
	s := strings.Repeat(".", n)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = s
	}
	col := make([]bool, n)
	pos := make([]bool, n*2-1)
	neg := make([]bool, n*2-1)
	var o [][]string
	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			t := make([]string, n)
			copy(t, a)
			o = append(o, t)
			return
		}
		for j := 0; j < n; j++ {
			if col[j] || pos[i-j+n-1] || neg[i+j] {
				continue
			}
			col[j] = true
			pos[i-j+n-1] = true
			neg[i+j] = true
			bak := a[i]
			c := []byte(a[i])
			c[j] = 'Q'
			a[i] = string(c)
			backtrack(i + 1)
			a[i] = bak
			col[j] = false
			pos[i-j+n-1] = false
			neg[i+j] = false
		}
	}
	backtrack(0)
	return o
}
