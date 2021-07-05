package main

// Tags:
// Union-Find

func findLatestStep(arr []int, m int) int {
	n := len(arr)
	uf := NewUnionFind(n + 1)
	on := make([]bool, n+1)
	cnt := map[int]int{}

	out := -1
	for j, i := range arr {
		on[i] = true
		cnt[uf.GetSize(i)]++
		if i-1 > 0 && on[i-1] {
			cnt[uf.GetSize(i)]--
			cnt[uf.GetSize(i-1)]--
			uf.Union(i, i-1)
			cnt[uf.GetSize(i)]++
		}
		if i+1 <= n && on[i+1] {
			cnt[uf.GetSize(i)]--
			cnt[uf.GetSize(i+1)]--
			uf.Union(i, i+1)
			cnt[uf.GetSize(i)]++
		}
		if cnt[m] > 0 {
			out = j + 1
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
	if x == y {
		return
	}
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}

func (uf UnionFind) GetSize(x int) int {
	return uf.size[uf.Find(x)]
}
