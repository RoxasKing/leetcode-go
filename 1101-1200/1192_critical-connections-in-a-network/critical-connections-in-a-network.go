package main

// Difficulty:
// Hard

// Tags:
// Biconnected Component
// DFS

func criticalConnections(n int, connections [][]int) [][]int {
	g := make([][]int, n)
	for _, e := range connections {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	idx := 0
	d := make([]int, n)
	low := make([]int, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
	var dfs func(int)
	dfs = func(u int) {
		idx++
		d[u] = idx
		low[u] = idx
		for _, v := range g[u] {
			if d[v] == 0 {
				parent[v] = u
				dfs(v)
				low[u] = min(low[u], low[v])
			} else if v != parent[u] {
				low[u] = min(low[u], low[v])
			}
		}
	}
	for u := 0; u < n; u++ {
		if d[u] == 0 {
			dfs(u)
		}
	}
	o := [][]int{}
	for v := 0; v < n; v++ {
		if u := parent[v]; u != -1 && d[v] == low[v] {
			o = append(o, []int{u, v})
		}
	}
	return o
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
