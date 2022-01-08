package main

import "runtime/debug"

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func cherryPickup(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][][2]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([][2]int, n)
		for j := i + 1; j < n; j++ {
			f[i][j][0] = -1
		}
	}
	f[0][n-1][0] = grid[0][0] + grid[0][n-1]
	cur, tmp := 0, 1
	for k := 1; k < m; k++ {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				f[i][j][tmp] = -1
				for a := Max(0, i-1); a <= Min(n-1, i+1); a++ {
					for b := Max(a+1, j-1); b <= Min(n-1, j+1); b++ {
						if f[a][b][cur] == -1 {
							continue
						}
						f[i][j][tmp] = Max(f[i][j][tmp], f[a][b][cur]+grid[k][i]+grid[k][j])
					}
				}
			}
		}
		cur, tmp = tmp, cur
	}
	out := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			out = Max(out, f[i][j][cur])
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

func init() { debug.SetGCPercent(-1) }
