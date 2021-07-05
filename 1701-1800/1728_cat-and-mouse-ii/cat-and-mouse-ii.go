package main

// Tags:
// BFS
// Hash

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	rows, cols := len(grid), len(grid[0])
	var x1, y1, x2, y2 int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'M' {
				x1, y1 = i, j
			} else if grid[i][j] == 'C' {
				x2, y2 = i, j
			}
		}
	}

	f := map[[5]int]int{}
	return dp(f, grid, catJump, mouseJump, rows, cols, x1, y1, x2, y2, 0) == 1
}

func dp(f map[[5]int]int, grid []string, catJump, mouseJump, m, n, x1, y1, x2, y2, k int) int {
	if k >= 70 {
		return 0
	}
	if val, ok := f[[5]int{x1, y1, x2, y2, k}]; ok {
		return val
	}
	if k%2 == 0 {
		for _, d := range direction {
			for j := 0; j <= mouseJump; j++ {
				x, y := x1+j*d[0], y1+j*d[1]
				if x < 0 || m-1 < x || y < 0 || n-1 < y || grid[x][y] == '#' {
					break
				} else if x == x2 && y == y2 {
					continue
				}
				if grid[x][y] == 'F' || dp(f, grid, catJump, mouseJump, m, n, x, y, x2, y2, k+1) == 1 {
					f[[5]int{x1, y1, x2, y2, k}] = 1
					return 1
				}
			}
		}
		f[[5]int{x1, y1, x2, y2, k}] = 0
		return 0
	} else {
		for _, d := range direction {
			for j := 0; j <= catJump; j++ {
				x, y := x2+j*d[0], y2+j*d[1]
				if x < 0 || m-1 < x || y < 0 || n-1 < y || grid[x][y] == '#' {
					break
				}
				if x == x1 && y == y1 || grid[x][y] == 'F' ||
					dp(f, grid, catJump, mouseJump, m, n, x1, y1, x, y, k+1) == 0 {
					f[[5]int{x1, y1, x2, y2, k}] = 0
					return 0
				}
			}
		}
		f[[5]int{x1, y1, x2, y2, k}] = 1
		return 1
	}
}

var direction = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
