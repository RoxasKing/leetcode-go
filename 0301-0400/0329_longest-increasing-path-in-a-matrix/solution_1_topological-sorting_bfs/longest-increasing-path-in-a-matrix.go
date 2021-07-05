package main

// Tags:
// Topological Sorting + BFS
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	size := m * n
	edges := make([][]int, size)
	indeg := make([]int, size)
	for i := 0; i < size; i++ {
		r, c := i/n, i%n
		for _, f := range forwards {
			nr, nc := r+f[0], c+f[1]
			if nr < 0 || m-1 < nr || nc < 0 || n-1 < nc || matrix[nr][nc] <= matrix[r][c] {
				continue
			}
			j := nr*n + nc
			edges[i] = append(edges[i], j)
			indeg[j]++
		}
	}

	q := [][2]int{}
	for i := 0; i < size; i++ {
		if indeg[i] == 0 {
			q = append(q, [2]int{i, 1})
		}
	}

	out := 0
	for len(q) > 0 {
		u, dist := q[0][0], q[0][1]
		q = q[1:]
		out = Max(out, dist)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, [2]int{v, dist + 1})
			}
		}
	}
	return out
}

var forwards = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
