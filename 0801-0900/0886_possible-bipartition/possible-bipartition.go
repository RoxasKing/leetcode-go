package main

// Difficulty:
// Medium

// Tags:
// Graph
// DFS

func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n+1)
	for _, p := range dislikes {
		g[p[0]] = append(g[p[0]], p[1])
		g[p[1]] = append(g[p[1]], p[0])
	}
	f := make([]int, n+1)
	var dfs func(x int) bool
	dfs = func(x int) bool {
		for _, y := range g[x] {
			if f[y] == 0 {
				f[y] = 0b11 ^ f[x]
				if !dfs(y) {
					return false
				}
			} else if f[x] == f[y] {
				return false
			}
		}
		return true
	}
	for i := 1; i <= n; i++ {
		if f[i] != 0 {
			continue
		}
		f[i] = 1
		if !dfs(i) {
			return false
		}
	}
	return true
}
