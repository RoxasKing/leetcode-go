package main

// Difficulty:
// Hard

// Tags:
// DFS
// Memorization
// Dynamic Programming

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	m, n := len(grid), len(grid[0])
	f := map[[5]int]bool{}
	var dp func(xc, yc, xm, ym, t int) bool
	dp = func(xc, yc, xm, ym, t int) bool {
		if t >= 68 {
			return false
		}
		key := [5]int{xc, yc, xm, ym, t}
		if v, ok := f[key]; ok {
			return v
		}
		x0, y0 := xc, yc
		jump := catJump
		if t&1 == 1 {
			x0, y0 = xm, ym
			jump = mouseJump
		}
		for i := 0; i < 4; i++ {
			for j := 0; j <= jump; j++ {
				x, y := x0+dirs[i]*j, y0+dirs[i+1]*j
				if x < 0 || m-1 < x || y < 0 || n-1 < y || grid[x][y] == '#' {
					break
				}
				if t&1 == 1 && xc == x && yc == y {
					continue
				}
				if grid[x][y] == 'F' ||
					t&1 == 0 && (x == xm && y == ym || !dp(x, y, xm, ym, t+1)) ||
					t&1 == 1 && dp(xc, yc, x, y, t+1) {
					f[key] = t&1 == 1
					return t&1 == 1
				}
			}
		}
		f[key] = (1 - t&1) == 1
		return (1 - t&1) == 1
	}
	var xc, yc, xm, ym int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'C' {
				xc, yc = i, j
			}
			if grid[i][j] == 'M' {
				xm, ym = i, j
			}
		}
	}
	return dp(xc, yc, xm, ym, 1)
}

var dirs = []int{-1, 0, 1, 0, -1}
