package main

// Difficulty:
// Hard

// Tags:
// DFS
// Memorization
// Dynamic Programming

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	m, n := len(grid), len(grid[0])
	f := map[[5]int]int{}
	var dp func(xc, yc, xm, ym, t int) int // 0: cat win ; 1: mouse win
	dp = func(xc, yc, xm, ym, t int) int {
		if t >= 68 {
			return 0
		}
		if v, ok := f[[5]int{xc, yc, xm, ym, t}]; ok {
			return v
		}
		x0, y0 := xc, yc
		jump := catJump
		if t&1 == 1 {
			x0, y0 = xm, ym
			jump = mouseJump
		}
		for k := 0; k < 4; k++ {
			for j := 0; j <= jump; j++ {
				x, y := x0+j*dirs[k], y0+j*dirs[k+1]
				if x < 0 || m-1 < x || y < 0 || n-1 < y || grid[x][y] == '#' {
					break
				}
				if t&1 == 1 && xc == x && yc == y {
					continue
				}
				if grid[x][y] == 'F' ||
					t&1 == 0 && x == xm && y == ym ||
					t&1 == 0 && dp(x, y, xm, ym, t+1) == 0 ||
					t&1 == 1 && dp(xc, yc, x, y, t+1) == 1 {
					f[[5]int{xc, yc, xm, ym, t}] = t & 1
					return t & 1
				}
			}
		}
		f[[5]int{xc, yc, xm, ym, t}] = 1 - t&1
		return 1 - t&1
	}
	var xc, yc, xm, ym int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'C' {
				xc, yc = i, j
			} else if grid[i][j] == 'M' {
				xm, ym = i, j
			}
		}
	}
	return dp(xc, yc, xm, ym, 1) == 1
}

var dirs = []int{-1, 0, 1, 0, -1}
