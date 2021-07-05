package main

// Tags:
// Union-Find
func minSwapsCouples(row []int) int {
	n := len(row) >> 1
	uf := newUnionFind(n)
	for i := 0; i < len(row); i += 2 {
		c1, c2 := row[i]>>1, row[i+1]>>1
		if c1 == c2 {
			continue
		}
		if uf.find(c1) == uf.find(c2) {
			continue
		}
		uf.union(c1, c2)
	}

	out := 0
	mark := make([]bool, n)
	for i := 0; i < n; i++ {
		x := uf.find(i)
		if mark[x] {
			out++
			continue
		}
		mark[x] = true
	}
	return out
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
