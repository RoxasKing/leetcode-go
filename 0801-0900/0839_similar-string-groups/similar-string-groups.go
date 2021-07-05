package main

// Tags:
// Union-Find
func numSimilarGroups(strs []string) int {
	n := len(strs)
	uf := newUnionFind(n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if isSimilar(strs[i], strs[j]) {
				uf.union(i, j)
			}
		}
	}
	mark := make([]bool, n)
	out := 0
	for i := range strs {
		x := uf.find(i)
		if !mark[x] {
			out++
			mark[x] = true
		}
	}
	return out
}

func isSimilar(s1, s2 string) bool {
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
			if count > 2 {
				return false
			}
		}
	}
	return true
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
