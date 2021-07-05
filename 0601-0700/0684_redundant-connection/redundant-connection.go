package main

// Tags:
// Union-Find
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := newUnionFind(n)
	for _, edge := range edges {
		ancestorX := uf.find(edge[0] - 1)
		ancestorY := uf.find(edge[1] - 1)
		if ancestorX != ancestorY {
			uf.union(ancestorX, ancestorY)
		} else {
			return edge
		}
	}
	return []int{}
}

type unionFind struct {
	ancestor []int
	isEnd    []bool
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := range ancestor {
		ancestor[i] = i
	}
	isEnd := make([]bool, n)
	return unionFind{ancestor: ancestor, isEnd: isEnd}
}

func (uf unionFind) find(x int) int {
	if uf.isEnd[uf.ancestor[x]] {
		return uf.ancestor[x]
	}
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
		uf.isEnd[x] = false
		uf.isEnd[uf.ancestor[x]] = true
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(x, y int) {
	ancestorX := uf.find(x)
	ancestorY := uf.find(y)
	uf.ancestor[ancestorX] = ancestorY
	uf.isEnd[ancestorX] = false
	uf.isEnd[ancestorY] = true
}
