package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Union-Find
// Sorting

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	uf := newUnionFind(n)
	for _, p := range pairs {
		uf.union(p[0], p[1])
	}
	h := map[int][]int{}
	for i := 0; i < n; i++ {
		x := uf.find(i)
		h[x] = append(h[x], i)
	}
	o := make([]byte, n)
	for _, a := range h {
		b := make([]int, len(a))
		copy(b, a)
		sort.Slice(a, func(i, j int) bool { return s[a[i]] < s[a[j]] })
		for i, k := range b {
			o[k] = s[a[i]]
		}
	}
	return string(o)
}

type unionFind struct{ parent, size []int }

func newUnionFind(n int) unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return unionFind{parent, size}
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) union(x, y int) {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return
	}
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}
