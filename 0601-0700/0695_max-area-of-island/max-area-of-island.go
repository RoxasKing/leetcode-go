package main

// Tags:
// Union-Find
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	uf := NewUnionFind(m * n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			x := i*n + j
			if i+1 < m && grid[i+1][j] == 1 {
				y := (i+1)*n + j
				uf.Union(x, y)
			}
			if j+1 < n && grid[i][j+1] == 1 {
				y := i*n + j + 1
				uf.Union(x, y)
			}
		}
	}

	out := 0
	cnt := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		r, c := i/n, i%n
		if grid[r][c] == 0 {
			continue
		}
		x := uf.Find(i)
		cnt[x]++
		if cnt[x] > out {
			out = cnt[x]
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
