package main

import "sort"

// Union-Find
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	m := len(edges)

	idxes := make([]int, m)
	for i := 0; i < m; i++ {
		idxes[i] = i
	}
	sort.Slice(idxes, func(i, j int) bool { return edges[idxes[i]][2] < edges[idxes[j]][2] })

	// first, find min weight MST
	minWeight := 0
	visited := make([]bool, m)
	minSet := []int{}
	others := []int{}
	uf := newUnionFind(n)
	for _, i := range idxes {
		edge := edges[i]
		x, y, weight := edge[0], edge[1], edge[2]
		if uf.find(x) == uf.find(y) {
			others = append(others, i)
			continue
		}
		minWeight += weight
		visited[i] = true
		minSet = append(minSet, i)
		uf.union(x, y)
	}

	// then, find critical edges
	critical := make([]bool, m)
	for _, i := range minSet {
		uf := newUnionFind(n)
		curWeight := 0
		for _, j := range idxes {
			if j == i {
				continue
			}
			edge := edges[j]
			x, y, weight := edge[0], edge[1], edge[2]
			if uf.find(x) == uf.find(y) {
				continue
			}
			uf.union(x, y)
			curWeight += weight
		}
		if curWeight != minWeight {
			critical[i] = true
		}
	}

	// last, find pseudo-critical edges
	for _, i := range others {
		edge := edges[i]
		x, y, weight := edge[0], edge[1], edge[2]
		uf := newUnionFind(n)
		curWeight := weight
		uf.union(x, y)
		for _, j := range minSet {
			edge := edges[j]
			x, y, weight := edge[0], edge[1], edge[2]
			if uf.find(x) == uf.find(y) {
				continue
			}
			curWeight += weight
			uf.union(x, y)
		}
		if curWeight == minWeight {
			visited[i] = true
		}
	}

	out := make([][]int, 2)
	for i := 0; i < m; i++ {
		if critical[i] {
			out[0] = append(out[0], i)
		} else if visited[i] {
			out[1] = append(out[1], i)
		}
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

	for i := range parent {
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
