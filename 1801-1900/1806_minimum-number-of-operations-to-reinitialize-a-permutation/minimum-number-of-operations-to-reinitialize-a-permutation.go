package main

// Tags:
// Union-Find
func reinitializePermutation(n int) int {
	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		x := i / 2
		if i&1 == 1 {
			x = n/2 + (i-1)/2
		}
		uf.Union(i, x)
	}

	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		cnt[uf.Find(i)]++
	}

	out := 0
	for i := 0; i < n; i++ {
		out = Max(out, cnt[i])
	}
	return out
}

type unionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return unionFind{
		parent: parent,
		size:   size,
	}
}

func (uf unionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)

	if x == y {
		return
	}

	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}

	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
