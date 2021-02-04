package main

import "sort"

/*
  You are a hiker preparing for an upcoming hike. You are given heights, a 2D array of size rows x columns, where heights[row][col] represents the height of cell (row, col). You are situated in the top-left cell, (0, 0), and you hope to travel to the bottom-right cell, (rows-1, columns-1) (i.e., 0-indexed). You can move up, down, left, or right, and you wish to find a route that requires the minimum effort.
  A route's effort is the maximum absolute difference in heights between two consecutive cells of the route.
  Return the minimum effort required to travel from the top-left cell to the bottom-right cell.

  Example 1:
    Input: heights = [[1,2,2],[3,8,2],[5,3,5]]
    Output: 2
    Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2 in consecutive cells.
    This is better than the route of [1,2,2,2,5], where the maximum absolute difference is 3.

  Example 2:
    Input: heights = [[1,2,3],[3,8,4],[5,3,5]]
    Output: 1
    Explanation: The route of [1,2,3,4,5] has a maximum absolute difference of 1 in consecutive cells, which is better than route [1,3,5,3,5].

  Example 3:
    Input: heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
    Output: 0
    Explanation: This route does not require any effort.

  Constraints:
    rows == heights.length
    columns == heights[i].length
    1 <= rows, columns <= 100
    1 <= heights[i][j] <= 106

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/path-with-minimum-effort
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	links := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				a, b, diff := (i-1)*n+j, i*n+j, Abs(heights[i-1][j]-heights[i][j])
				links = append(links, []int{a, b, diff})
			}
			if j > 0 {
				a, b, diff := i*n+j-1, i*n+j, Abs(heights[i][j-1]-heights[i][j])
				links = append(links, []int{a, b, diff})
			}
		}
	}
	sort.Slice(links, func(i, j int) bool { return links[i][2] < links[j][2] })
	out := 0
	uf := newUnionFind(m * n)
	for _, link := range links {
		x, y, diff := link[0], link[1], link[2]
		if uf.find(x) == uf.find(y) {
			continue
		}
		uf.union(x, y)
		out = Max(out, diff)
		if uf.find(0) == uf.find(m*n-1) {
			break
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
	if uf.size[x] > uf.size[y] {
		x, y = y, x
	}
	uf.parent[x] = y
	uf.size[y] += uf.size[x]
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
