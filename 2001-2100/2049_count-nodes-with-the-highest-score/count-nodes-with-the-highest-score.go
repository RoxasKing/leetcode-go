package main

// Difficulty:
// Medium

// Tags:
// Graph
// DFS

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	g := make([][]int, n)
	for i, p := range parents {
		if p == -1 {
			continue
		}
		g[p] = append(g[p], i)
	}
	c := make([]int, n)
	var dfs func(int)
	dfs = func(u int) {
		c[u] = 1
		for _, v := range g[u] {
			dfs(v)
			c[u] += c[v]
		}
	}
	dfs(0)
	o, x := 1, 1
	for _, v := range g[0] {
		x *= c[v]
	}
	for u := 1; u < n; u++ {
		t := c[0] - c[u]
		for _, v := range g[u] {
			t *= c[v]
		}
		if x == t {
			o++
		} else if x < t {
			o, x = 1, t
		}
	}
	return o
}
