package main

// Difficulty:
// Hard

// Tags:
// BFS
// Bitmask

func shortestPathAllKeys(grid []string) int {
	dirs := []int{-1, 0, 1, 0, -1}
	m, n := len(grid), len(grid[0])
	h, keys := [26]int{}, []rune{}
	var x0, y0 int
	for i, rs := range grid {
		for j, c := range rs {
			if 'a' <= c && c <= 'z' {
				h[c-'a'], keys = len(keys), append(keys, c)
			} else if c == '@' {
				x0, y0 = i, j
			}
		}
	}
	f := make([][][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = make([]int, 1<<len(keys))
			for k := 0; k < 1<<len(keys); k++ {
				f[i][j][k] = -1
			}
		}
	}
	f[x0][y0][0] = 0
	for q := [][]int{{x0, y0, 0}}; len(q) > 0; q = q[1:] {
		p := q[0]
		i, j, mask := p[0], p[1], p[2]
		for k := 0; k < 4; k++ {
			x, y := i+dirs[k], j+dirs[k+1]
			if x < 0 || m-1 < x || y < 0 || n-1 < y || grid[x][y] == '#' {
				continue
			}
			if c := rune(grid[x][y]); c == '.' || c == '@' {
				if f[x][y][mask] == -1 {
					f[x][y][mask] = f[i][j][mask] + 1
					q = append(q, []int{x, y, mask})
				}
			} else if 'a' <= c && c <= 'z' {
				if t := mask | 1<<h[c-'a']; f[x][y][t] == -1 {
					f[x][y][t] = f[i][j][mask] + 1
					if t == 1<<len(keys)-1 {
						return f[x][y][t]
					}
					q = append(q, []int{x, y, t})
				}
			} else {
				if id := h[c-'A']; mask>>id&1 == 1 && f[x][y][mask] == -1 {
					f[x][y][mask] = f[i][j][mask] + 1
					q = append(q, []int{x, y, mask})
				}
			}
		}
	}
	return -1
}
