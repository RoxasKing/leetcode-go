package main

// Tags:
// Union-Find
func makeConnected(n int, connections [][]int) int {
	uf := newUnionFind(n)
	cables := 0
	for _, conn := range connections {
		x, y := conn[0], conn[1]
		if uf.find(x) == uf.find(y) {
			cables++
			continue
		}
		uf.union(x, y)
	}
	mark := make([]bool, n)
	group := -1
	for i := 0; i < n; i++ {
		x := uf.find(i)
		if !mark[x] {
			mark[x] = true
			group++
		}
	}
	if group > cables {
		return -1
	}
	return group
}

type unionFind struct {
	parent []int
	size   []int
}

func newUnionFind(n int) unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return unionFind{parent: parent, size: size}
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) union(x, y int) {
	x = uf.find(x)
	y = uf.find(y)
	if x == y {
		return
	}
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}

	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}
