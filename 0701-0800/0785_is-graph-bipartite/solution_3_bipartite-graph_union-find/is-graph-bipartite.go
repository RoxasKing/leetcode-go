package main

// Difficulty:
// Medium

// Tags:
// Bipartite Graph
// Union-Find

func isBipartite(graph [][]int) bool {
	n := len(graph)
	uf := newUnionFind(n)
	for u, vs := range graph {
		if len(vs) == 0 {
			continue
		}
		for _, v := range vs {
			if uf.find(u) == uf.find(v) {
				return false
			}
		}
		x := vs[0]
		for _, y := range vs[1:] {
			uf.union(x, y)
		}
	}
	return true
}

type unionFind struct{ parent, size []int }

func newUnionFind(n int) unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return unionFind{parent, size}
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) union(x, y int) {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return
	}
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}
