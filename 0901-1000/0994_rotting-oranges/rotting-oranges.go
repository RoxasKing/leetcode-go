package main

// Difficulty:
// Medium

// Tags:
// BFS

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	q := [][]int{}
	remain := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				remain++
			}
		}
	}
	cnt := 0
	for len(q) > 0 && remain > 0 {
		cnt++
		k := len(q)
		for _, e := range q {
			x, y := e[0], e[1]
			for i := 0; i < 4; i++ {
				a, b := x+rs[i], y+cs[i]
				if a < 0 || m-1 < a || b < 0 || n-1 < b || grid[a][b] != 1 {
					continue
				}
				remain--
				grid[a][b] = 2
				q = append(q, []int{a, b})
			}
		}
		q = q[k:]
	}
	if remain > 0 {
		return -1
	}
	return cnt
}

var rs = []int{-1, 1, 0, 0}
var cs = []int{0, 0, -1, 1}
