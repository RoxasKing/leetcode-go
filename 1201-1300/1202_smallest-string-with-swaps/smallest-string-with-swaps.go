package main

import (
	"sort"
)

// Tags:
// Union-Find
func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)

	uf := newUnionFind(n)
	for _, pair := range pairs {
		uf.union(pair[0], pair[1])
	}

	groups := make(map[int][]int)
	for i := 0; i < n; i++ {
		a := uf.find(i)
		groups[a] = append(groups[a], i)
	}

	bs := []byte(s)
	for _, group := range groups {
		if len(group) < 2 {
			continue
		}
		arr := make([]byte, len(group))
		for i := range group {
			arr[i] = bs[group[i]]
		}
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		for i := range arr {
			bs[group[i]] = arr[i]
		}
	}

	return string(bs)
}

type unionFind struct {
	ancestor []int
	isEnd    []bool
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	isEnd := make([]bool, n)
	for i := range ancestor {
		ancestor[i] = i
	}
	return unionFind{ancestor: ancestor, isEnd: isEnd}
}

func (uf unionFind) find(x int) int {
	if uf.isEnd[uf.ancestor[x]] {
		return uf.ancestor[x]
	}

	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
	}
	uf.isEnd[x] = false
	uf.isEnd[uf.ancestor[x]] = true

	return uf.ancestor[x]
}

func (uf unionFind) union(x, y int) {
	xAncestor := uf.find(x)
	yAncestor := uf.find(y)
	uf.ancestor[xAncestor] = yAncestor
	uf.isEnd[xAncestor] = false
	uf.isEnd[yAncestor] = true
}
