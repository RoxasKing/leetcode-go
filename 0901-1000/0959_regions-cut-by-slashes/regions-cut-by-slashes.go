package main

// Tags:
// Union-Find
func regionsBySlashes(grid []string) int {
	n := len(grid)
	size := 4 * n * n
	uf := newUnionFind(size)
	for i := range grid {
		for j := range grid[i] {
			start := 4 * (n*j + i)
			up, right, down, left := start, start+1, start+2, start+3
			if grid[i][j] == '/' {
				uf.union(right, down)
				uf.union(up, left)
			} else if grid[i][j] == '\\' {
				uf.union(up, right)
				uf.union(down, left)
			} else {
				uf.union(right, up)
				uf.union(down, up)
				uf.union(left, up)
			}
			if i+1 < n {
				rightStart := start + 4
				rightLeft := rightStart + 3
				uf.union(rightLeft, right)
			}
			if j+1 < n {
				downStart := start + 4*n
				downUp := downStart
				uf.union(downUp, down)
			}
		}
	}
	out := 0
	mark := make([]bool, size)
	for i := 0; i < size; i++ {
		x := uf.find(i)
		if !mark[x] {
			out++
			mark[x] = true
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
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}
