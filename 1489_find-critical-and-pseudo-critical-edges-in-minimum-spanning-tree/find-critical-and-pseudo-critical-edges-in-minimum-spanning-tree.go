package main

import "sort"

/*
  Given a weighted undirected connected graph with n vertices numbered from 0 to n - 1, and an array edges where edges[i] = [ai, bi, weighti] represents a bidirectional and weighted edge between nodes ai and bi. A minimum spanning tree (MST) is a subset of the graph's edges that connects all vertices without cycles and with the minimum possible total edge weight.

  Find all the critical and pseudo-critical edges in the given graph's minimum spanning tree (MST). An MST edge whose deletion from the graph would cause the MST weight to increase is called a critical edge. On the other hand, a pseudo-critical edge is that which can appear in some MSTs but not all.

  Note that you can return the indices of the edges in any order.

  Example 1:
    Input: n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]
    Output: [[0,1],[2,3,4,5]]
    Explanation: The figure above describes the graph.
    The following figure shows all the possible MSTs:
    Notice that the two edges 0 and 1 appear in all MSTs, therefore they are critical edges, so we return them in the first list of the output.
    The edges 2, 3, 4, and 5 are only part of some MSTs, therefore they are considered pseudo-critical edges. We add them to the second list of the output.

  Example 2:
    Input: n = 4, edges = [[0,1,1],[1,2,1],[2,3,1],[0,3,1]]
    Output: [[],[0,1,2,3]]
    Explanation: We can observe that since all 4 edges have equal weight, choosing any 3 edges from the given 4 will yield an MST. Therefore all 4 edges are pseudo-critical.


  Constraints:
    2 <= n <= 100
    1 <= edges.length <= min(200, n * (n - 1) / 2)
    edges[i].length == 3
    0 <= ai < bi < n
    1 <= weighti <= 1000
    All pairs (ai, bi) are distinct.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	m := len(edges)

	idxes := make([]int, m)
	for i := 0; i < m; i++ {
		idxes[i] = i
	}
	sort.Slice(idxes, func(i, j int) bool { return edges[idxes[i]][2] < edges[idxes[j]][2] })

	// first, find min weight MST
	minWeight := 0
	visited := make([]bool, m)
	minSet := []int{}
	others := []int{}
	uf := newUnionFind(n)
	for _, i := range idxes {
		edge := edges[i]
		x, y, weight := edge[0], edge[1], edge[2]
		if uf.find(x) == uf.find(y) {
			others = append(others, i)
			continue
		}
		minWeight += weight
		visited[i] = true
		minSet = append(minSet, i)
		uf.union(x, y)
	}

	// then, find critical edges
	critical := make([]bool, m)
	for _, i := range minSet {
		uf := newUnionFind(n)
		curWeight := 0
		for _, j := range idxes {
			if j == i {
				continue
			}
			edge := edges[j]
			x, y, weight := edge[0], edge[1], edge[2]
			if uf.find(x) == uf.find(y) {
				continue
			}
			uf.union(x, y)
			curWeight += weight
		}
		if curWeight != minWeight {
			critical[i] = true
		}
	}

	// last, find pseudo-critical edges
	for _, i := range others {
		edge := edges[i]
		x, y, weight := edge[0], edge[1], edge[2]
		uf := newUnionFind(n)
		curWeight := weight
		uf.union(x, y)
		for _, j := range minSet {
			edge := edges[j]
			x, y, weight := edge[0], edge[1], edge[2]
			if uf.find(x) == uf.find(y) {
				continue
			}
			curWeight += weight
			uf.union(x, y)
		}
		if curWeight == minWeight {
			visited[i] = true
		}
	}

	out := make([][]int, 2)
	for i := 0; i < m; i++ {
		if critical[i] {
			out[0] = append(out[0], i)
		} else if visited[i] {
			out[1] = append(out[1], i)
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

	for i := range parent {
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
