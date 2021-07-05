package main

import "sort"

// Prim's Algorithm + Greedy Algorithm
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)

	for i := 1; i < n; i++ {
		minDist[i] = 1<<31 - 1
	}

	out := 0

	for i := 0; i < n; i++ {
		next := 0
		dist := 1<<31 - 1
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < dist {
				next = j
				dist = minDist[j]
			}
		}

		out += dist
		visited[next] = true

		for j := 0; j < n; j++ {
			if !visited[j] {
				minDist[j] = Min(minDist[j], Abs(points[j][0]-points[next][0])+Abs(points[j][1]-points[next][1]))
			}
		}
	}

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Union-Find
func minCostConnectPoints2(points [][]int) int {
	n := len(points)

	dists := []distance{}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			val := Abs(points[i][0]-points[j][0]) + Abs(points[i][1]-points[j][1])
			dists = append(dists, distance{a: i, b: j, val: val})
		}
	}

	sort.Slice(dists, func(i, j int) bool { return dists[i].val < dists[j].val })

	uf := newUnionFind(n)
	out := 0

	for _, dist := range dists {
		x, y, val := dist.a, dist.b, dist.val
		if uf.find(x) == uf.find(y) {
			continue
		}
		uf.union(x, y)
		out += val
	}

	return out
}

type distance struct {
	a, b int
	val  int
}

type unionFind struct {
	ancestor []int
	isEnd    []bool
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := 0; i < n; i++ {
		ancestor[i] = i
	}
	isEnd := make([]bool, n)
	return unionFind{ancestor: ancestor, isEnd: isEnd}
}

func (uf unionFind) find(x int) int {
	if uf.isEnd[uf.ancestor[x]] {
		return uf.ancestor[x]
	}
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
		uf.isEnd[x] = false
		uf.isEnd[uf.ancestor[x]] = true
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(x, y int) {
	ancestorX := uf.find(x)
	ancestorY := uf.find(y)
	uf.ancestor[ancestorX] = ancestorY
	uf.isEnd[ancestorX] = false
	uf.isEnd[ancestorY] = true
}
