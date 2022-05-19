package main

// Difficulty:
// Hard

// Tags:
// DFS
// Dynamic Programming

func longestIncreasingPath(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}

	var out int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(mat, f, m, n, i, j)
			out = Max(out, f[i][j])
		}
	}
	return out
}

func dfs(mat, f [][]int, m, n, i, j int) {
	if f[i][j] != 0 {
		return
	}
	f[i][j] = 1
	for k := 0; k < 4; k++ {
		x, y := i+r[k], j+c[k]
		if x < 0 || m-1 < x || y < 0 || n-1 < y || mat[x][y] <= mat[i][j] {
			continue
		}
		dfs(mat, f, m, n, x, y)
		f[i][j] = Max(f[i][j], f[x][y]+1)
	}
}

var r = []int{-1, 1, 0, 0}
var c = []int{0, 0, -1, 1}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
