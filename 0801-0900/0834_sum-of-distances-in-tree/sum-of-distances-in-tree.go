package main

// Tags:
// Dynamic Programming
// DFS

func sumOfDistancesInTree(n int, edges [][]int) []int {
	g = make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	c = make([]int, n)
	f = make([]int, n)
	out = make([]int, n)

	dfs1(0, -1)
	dfs2(0, -1)

	return out
}

var g [][]int
var c, f, out []int

func dfs1(u, x int) {
	c[u] = 1
	for _, v := range g[u] {
		if v == x {
			continue
		}
		dfs1(v, u)
		c[u] += c[v]
		f[u] += f[v] + c[v]
	}
}

func dfs2(u, x int) {
	out[u] = f[u]
	for _, v := range g[u] {
		if v == x {
			continue
		}

		cu, fu := c[u], f[u]
		cv, fv := c[v], f[v]

		c[u] -= c[v]
		f[u] -= f[v] + c[v]
		c[v] += c[u]
		f[v] += f[u] + c[u]

		dfs2(v, u)

		c[u], f[u] = cu, fu
		c[v], f[v] = cv, fv
	}
}
