package main

import "sort"

// Difficult:
// Medium

// Tags:
// Union-Find

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	dists := []dist{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			v := Abs(points[i][0]-points[j][0]) + Abs(points[i][1]-points[j][1])
			dists = append(dists, dist{i, j, v})
		}
	}
	sort.Slice(dists, func(i, j int) bool { return dists[i].v < dists[j].v })
	uf := newUnionFind(n)
	out := 0
	for _, dist := range dists {
		a, b := dist.a, dist.b
		if uf.find(a) == uf.find(b) {
			continue
		}
		uf.union(a, b)
		out += dist.v
	}
	return out
}

type dist struct{ a, b, v int }

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
