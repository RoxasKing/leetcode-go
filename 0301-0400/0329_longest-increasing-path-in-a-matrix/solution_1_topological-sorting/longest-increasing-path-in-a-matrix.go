package main

// Difficulty:
// Hard

// Tags:
// Topological Sort
// BFS

func longestIncreasingPath(matrix [][]int) int {
	dirs := []int{-1, 0, 1, 0, -1}
	m, n := len(matrix), len(matrix[0])
	g := make([][]int, m*n)
	d := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			u := i*n + j
			for k := 0; k < 4; k++ {
				x, y := i+dirs[k], j+dirs[k+1]
				if x < 0 || m-1 < x || y < 0 || n-1 < y || matrix[x][y] <= matrix[i][j] {
					continue
				}
				v := x*n + y
				g[u] = append(g[u], v)
				d[v]++
			}
		}
	}
	q := []int{}
	for i := 0; i < m*n; i++ {
		if d[i] == 0 {
			q = append(q, i)
		}
	}
	o := 0
	for len(q) > 0 {
		o++
		k := len(q)
		for _, u := range q {
			for _, v := range g[u] {
				if d[v]--; d[v] == 0 {
					q = append(q, v)
				}
			}
		}
		q = q[k:]
	}
	return o
}
