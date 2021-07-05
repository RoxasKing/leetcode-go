package main

// Tags:
// DFS
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	size := m * n
	edges := make([][]int, size)
	for i := 0; i < size; i++ {
		r, c := i/n, i%n
		for _, f := range forwards {
			nr, nc := r+f[0], c+f[1]
			if nr < 0 || m-1 < nr || nc < 0 || n-1 < nc || matrix[nr][nc] <= matrix[r][c] {
				continue
			}
			j := nr*n + nc
			edges[i] = append(edges[i], j)
		}
	}

	f := make([]int, size)
	for i := 0; i < size; i++ {
		dfs(edges, f, i, 1)
	}

	out := 0
	for i := 0; i < size; i++ {
		out = Max(out, f[i])
	}
	return out
}

func dfs(edges [][]int, f []int, u, dist int) {
	if f[u] >= dist {
		return
	}
	f[u] = dist
	for _, v := range edges[u] {
		dfs(edges, f, v, dist+1)
	}
}

var forwards = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
