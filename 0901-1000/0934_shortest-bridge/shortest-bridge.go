package main

// Difficulty:
// Medium

// Tags:
// Graph
// BFS

func shortestBridge(grid [][]int) int {
	n := len(grid)
	dirs := []int{-1, 0, 1, 0, -1}
	var a, b [][]int
	for i, vs := range grid {
		for j, v := range vs {
			if v != 1 {
				continue
			}
			t := [][]int{}
			grid[i][j] = -1
			for q := [][]int{{i, j}}; len(q) > 0; q = q[1:] {
				p := q[0]
				t = append(t, p)
				for k := 0; k < 4; k++ {
					x, y := p[0]+dirs[k], p[1]+dirs[k+1]
					if x < 0 || n-1 < x || y < 0 || n-1 < y || grid[x][y] != 1 {
						continue
					}
					grid[x][y] = -1
					q = append(q, []int{x, y})
				}
			}
			if len(a) == 0 {
				a = t
			} else {
				b = t
			}
		}
	}
	o := 1<<31 - 1
	for _, p := range a {
		for _, q := range b {
			o = min(o, abs(p[0]-q[0])+abs(p[1]-q[1])-1)
		}
	}
	return o
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
