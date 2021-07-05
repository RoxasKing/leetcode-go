package main

import "sort"

// Union-Find
func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	links := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				a, b, diff := (i-1)*n+j, i*n+j, Abs(heights[i-1][j]-heights[i][j])
				links = append(links, []int{a, b, diff})
			}
			if j > 0 {
				a, b, diff := i*n+j-1, i*n+j, Abs(heights[i][j-1]-heights[i][j])
				links = append(links, []int{a, b, diff})
			}
		}
	}
	sort.Slice(links, func(i, j int) bool { return links[i][2] < links[j][2] })
	out := 0
	uf := newUnionFind(m * n)
	for _, link := range links {
		x, y, diff := link[0], link[1], link[2]
		if uf.find(x) == uf.find(y) {
			continue
		}
		uf.union(x, y)
		out = Max(out, diff)
		if uf.find(0) == uf.find(m*n-1) {
			break
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
	if uf.size[x] > uf.size[y] {
		x, y = y, x
	}
	uf.parent[x] = y
	uf.size[y] += uf.size[x]
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
