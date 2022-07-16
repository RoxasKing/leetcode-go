package main

// Difficulty:
// Medium

// Tags:
// Union-Find

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	area := m * n
	parent := make([]int, area)
	size := make([]int, area)
	for i := 0; i < area; i++ {
		parent[i] = i
		size[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x, y int) {
		x, y = find(x), find(y)
		if x == y {
			return
		}
		if size[x] < size[y] {
			x, y = y, x
		}
		parent[y] = x
		size[x] += size[y]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			u := i*n + j
			if i+1 < m && grid[i+1][j] == 1 {
				v := (i+1)*n + j
				union(u, v)
			}
			if j+1 < n && grid[i][j+1] == 1 {
				v := i*n + j + 1
				union(u, v)
			}
		}
	}
	o := 0
	cnt := make([]int, area)
	for i := 0; i < area; i++ {
		if grid[i/n][i%n] == 0 {
			continue
		}
		x := find(i)
		if cnt[x]++; o < cnt[x] {
			o = cnt[x]
		}
	}
	return o
}
