package main

// Difficulty:
// Hard

// Tags:
// Topological Sort
// BFS

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	g := make([][]int, m*n)
	d := make([]int, m*n)
	add := func(x1, y1, x2, y2 int) {
		u, v := x1*n+y1, x2*n+y2
		if matrix[x1][y1] > matrix[x2][y2] {
			u, v = v, u
		}
		g[u] = append(g[u], v)
		d[v]++
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i+1 < m && matrix[i][j] != matrix[i+1][j] {
				add(i, j, i+1, j)
			}
			if j+1 < n && matrix[i][j] != matrix[i][j+1] {
				add(i, j, i, j+1)
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
