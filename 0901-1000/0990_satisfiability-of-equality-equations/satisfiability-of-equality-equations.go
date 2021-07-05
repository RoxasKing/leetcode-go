package main

// Tags:
// Union-Find
func equationsPossible(equations []string) bool {
	uf := NewUnionFind(26)
	for _, e := range equations {
		x, y := int(e[0]-'a'), int(e[3]-'a')
		if e[1] == '=' {
			uf.Union(x, y)
		}
	}
	for _, e := range equations {
		x, y := int(e[0]-'a'), int(e[3]-'a')
		x, y = uf.Find(x), uf.Find(y)
		if e[1] == '!' && x == y {
			return false
		}
	}
	return true
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
	return UnionFind{
		parent: parent,
		size:   size,
	}
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
	if x == y {
		return
	}
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}
