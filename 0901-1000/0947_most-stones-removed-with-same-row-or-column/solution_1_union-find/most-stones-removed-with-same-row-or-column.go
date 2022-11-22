package main

// Difficulty:
// Medium

// Tags:
// Union-Find

func removeStones(stones [][]int) int {
	n := len(stones)
	parent := make([]int, n)
	size := make([]int, n)
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
	rs := [1e4 + 1][]int{}
	cs := [1e4 + 1][]int{}
	for i, e := range stones {
		rs[e[0]] = append(rs[e[0]], i)
		cs[e[1]] = append(cs[e[1]], i)
	}
	for _, a := range rs {
		for i := 1; i < len(a); i++ {
			union(a[i-1], a[i])
		}
	}
	for _, a := range cs {
		for i := 1; i < len(a); i++ {
			union(a[i-1], a[i])
		}
	}
	o, vis := n, [1e4 + 1]bool{}
	for i := 0; i < n; i++ {
		if p := find(i); !vis[p] {
			o--
			vis[p] = true
		}
	}
	return o
}
