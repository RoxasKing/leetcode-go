package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Union-Find

func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	uf := NewUnionFind(n)
	emailMap := map[string]int{}
	for x, arr := range accounts {
		for _, email := range arr[1:] {
			if y, ok := emailMap[email]; ok {
				uf.Union(x, y)
			} else {
				emailMap[email] = x
			}
		}
	}
	accountMap := map[int][]string{}
	for i := range accounts {
		x := uf.Find(i)
		if _, ok := accountMap[x]; !ok {
			accountMap[x] = append(accountMap[x], accounts[i][0])
		}
	}
	for email, x := range emailMap {
		x = uf.Find(x)
		accountMap[x] = append(accountMap[x], email)
	}
	out := make([][]string, 0, len(accountMap))
	for _, a := range accountMap {
		sort.Strings(a[1:])
		out = append(out, a)
	}
	return out
}

type UnionFind struct {
	parent, size []int
}

func NewUnionFind(n int) UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return UnionFind{parent: parent, size: size}
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
	uf.size[x] += uf.parent[y]
}
