package main

// Tags:
// Topological Sort
// BFS

func longestIncreasingPath(mat [][]int) int {
	r := [4]int{-1, 1, 0, 0}
	c := [4]int{0, 0, -1, 1}
	m, n := len(mat), len(mat[0])
	size := m * n
	g := make([][]int, size)
	d := make([]int, size)
	for u := 0; u < size; u++ {
		i, j := u/n, u%n
		for k := 0; k < 4; k++ {
			x, y := i+r[k], j+c[k]
			if x < 0 || m-1 < x || y < 0 || n-1 < y || mat[x][y] <= mat[i][j] {
				continue
			}
			v := x*n + y
			g[u] = append(g[u], v)
			d[v]++
		}
	}

	q := []int{}
	for u := 0; u < size; u++ {
		if d[u] == 0 {
			q = append(q, u)
		}
	}

	out := 0
	for len(q) > 0 {
		out++
		k := len(q)
		for _, u := range q {
			for _, v := range g[u] {
				d[v]--
				if d[v] == 0 {
					q = append(q, v)
				}
			}
		}
		q = q[k:]
	}
	return out
}
