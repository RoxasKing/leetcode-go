package main

import "sort"

/*
  You are given an array points representing integer coordinates of some points on a 2D-plane, where points[i] = [xi, yi].
  The cost of connecting two points [xi, yi] and [xj, yj] is the manhattan distance between them: |xi - xj| + |yi - yj|, where |val| denotes the absolute value of val.
  Return the minimum cost to make all points connected. All points are connected if there is exactly one simple path between any two points.

  Example 1:
    Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
    Output: 20
    Explanation:
      We can connect the points as shown above to get the minimum cost of 20.
      Notice that there is a unique path between every pair of points.

  Example 2:
    Input: points = [[3,12],[-2,5],[-4,1]]
    Output: 18

  Example 3:
    Input: points = [[0,0],[1,1],[1,0],[-1,1]]
    Output: 4

  Example 4:
    Input: points = [[-1000000,-1000000],[1000000,1000000]]
    Output: 4000000

  Example 5:
    Input: points = [[0,0]]
    Output: 0

  Constraints:
    1 <= points.length <= 1000
    -10^6 <= xi, yi <= 10^6
    All pairs (xi, yi) are distinct.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/min-cost-to-connect-all-points
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
