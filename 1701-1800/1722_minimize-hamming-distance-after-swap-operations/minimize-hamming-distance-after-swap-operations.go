package main

// Tags:
// Union-Find

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	for i := range source {
		source[i]--
		target[i]--
	}
	n := len(source)
	uf := NewUnionFind(n)
	for _, e := range allowedSwaps {
		u, v := e[0], e[1]
		uf.Union(u, v)
	}
	g := map[int][]int{}
	for i := 0; i < n; i++ {
		x := uf.Find(i)
		g[x] = append(g[x], i)
	}
	out := 0
	for _, arr := range g {
		cnt := map[int]int{}
		for _, id := range arr {
			cnt[source[id]]++
		}
		for _, id := range arr {
			cnt[target[id]]--
			if cnt[target[id]] < 0 {
				out++
			}
		}
	}
	return out
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return UnionFind{parent: parent, size: size}
}

func (uf UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf UnionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}
