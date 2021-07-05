package main

// Tags:
// Union-Find
func maxNumEdgesToRemove(n int, edges [][]int) int {
	uf1 := newUnionFind(n)
	uf2 := newUnionFind(n)
	count1 := 0
	count2 := 0
	remain1 := 0
	remain2 := 0
	remain3 := 0

	for _, e := range edges {
		typ, x, y := e[0], e[1]-1, e[2]-1
		if typ != 3 {
			continue
		}
		if uf1.find(x) == uf1.find(y) && uf2.find(x) == uf2.find(y) {
			remain3++
			continue
		}
		if uf1.find(x) != uf1.find(y) {
			uf1.union(x, y)
			count1++
		}
		if uf2.find(x) != uf2.find(y) {
			uf2.union(x, y)
			count2++
		}
	}

	for _, e := range edges {
		typ, x, y := e[0], e[1]-1, e[2]-1
		if typ != 1 {
			continue
		}
		if uf1.find(x) == uf1.find(y) {
			remain1++
			continue
		}
		uf1.union(x, y)
		count1++
	}

	for _, e := range edges {
		typ, x, y := e[0], e[1]-1, e[2]-1
		if typ != 2 {
			continue
		}
		if uf2.find(x) == uf2.find(y) {
			remain2++
			continue
		}
		uf2.union(x, y)
		count2++
	}

	if count1 < n-1 || count2 < n-1 {
		return -1
	}

	return remain1 + remain2 + remain3
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
	if uf.size[x] > uf.size[y] {
		x, y = y, x
	}
	uf.parent[x] = y
	uf.size[y] += uf.size[x]
}
