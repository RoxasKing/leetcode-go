package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Union-Find
// Sorting

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	a := []edge{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x := i*n + j
			if i+1 < m {
				y := (i+1)*n + j
				a = append(a, edge{x, y, abs(heights[i][j] - heights[i+1][j])})
			}
			if j+1 < n {
				y := i*n + j + 1
				a = append(a, edge{x, y, abs(heights[i][j] - heights[i][j+1])})
			}
		}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].d < a[j].d })
	o := 0
	uf := newUnionFind(m * n)
	for _, e := range a {
		x, y := e.x, e.y
		if uf.find(x) == uf.find(y) {
			continue
		}
		uf.union(x, y)
		if o < e.d {
			o = e.d
		}
		if uf.find(0) == uf.find(m*n-1) {
			break
		}
	}
	return o
}

type edge struct{ x, y, d int }

func abs(a int) int {
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
