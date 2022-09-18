package main

// Difficulty:
// Hard

// Tags:
// Union-Find

func largestIsland(grid [][]int) int {
	n := len(grid)
	parent := make([]int, n*n)
	size := make([]int, n*n)
	for i := range parent {
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
	for i, vs := range grid {
		for j, v := range vs {
			if v == 0 {
				continue
			}
			u := i*n + j
			if i+1 < n && grid[i+1][j] == 1 {
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
	dirs := []int{-1, 0, 1, 0, -1}
	for i, vs := range grid {
		for j, v := range vs {
			if v == 1 {
				o = max(o, size[find(i*n+j)])
				continue
			}
			vis := map[int]bool{}
			area := 1
			for k := 0; k < 4; k++ {
				x, y := i+dirs[k], j+dirs[k+1]
				if x < 0 || n-1 < x || y < 0 || n-1 < y || grid[x][y] == 0 {
					continue
				}
				id := find(x*n + y)
				if vis[id] {
					continue
				}
				vis[id] = true
				area += size[id]
			}
			o = max(o, area)
		}
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
