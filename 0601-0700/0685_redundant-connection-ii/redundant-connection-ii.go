package main

// Tags:
// Union-Find
func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	uf := newUnionFind(n + 1)
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}

	var conflictEdge, cycleEdge []int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if parent[to] != to { // to have two parent
			conflictEdge = edge
		} else {
			parent[to] = from
			if uf.find(from) == uf.find(to) { // from and to have unioned
				cycleEdge = edge
			} else {
				uf.union(from, to)
			}
		}
	}

	if conflictEdge == nil {
		return cycleEdge
	}

	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}

	return conflictEdge
}

type unionFind struct {
	ancestor []int
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := 0; i < n; i++ {
		ancestor[i] = i
	}
	return unionFind{ancestor}
}

func (uf unionFind) find(x int) int {
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(from, to int) {
	uf.ancestor[uf.find(from)] = uf.find(to)
}
