package main

// Difficulty:
// Hard

// Tags:
// Graph
// Hash

func checkWays(pairs [][]int) int {
	g := map[int]map[int]struct{}{}
	for _, p := range pairs {
		u, v := p[0], p[1]
		if _, ok := g[u]; !ok {
			g[u] = map[int]struct{}{}
		}
		if _, ok := g[v]; !ok {
			g[v] = map[int]struct{}{}
		}
		g[u][v] = struct{}{}
		g[v][u] = struct{}{}
	}
	root := -1
	for u := range g {
		if len(g[u]) == len(g)-1 {
			root = u
			break
		}
	}
	if root == -1 {
		return 0
	}
	out := 1
	for u := range g {
		if u == root {
			continue
		}
		d := len(g[u])
		p := -1
		pd := 1<<31 - 1
		for v := range g[u] {
			if len(g[v]) < pd && len(g[v]) >= d {
				p = v
				pd = len(g[v])
			}
		}
		if p == -1 {
			return 0
		}
		for v := range g[u] {
			if v == p {
				continue
			}
			if _, ok := g[p][v]; !ok {
				return 0
			}
		}
		if d == pd {
			out = 2
		}
	}
	return out
}
