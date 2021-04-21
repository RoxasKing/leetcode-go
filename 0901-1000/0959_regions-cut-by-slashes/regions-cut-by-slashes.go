package main

/*
  In a N x N grid composed of 1 x 1 squares, each 1 x 1 square consists of a /, \, or blank space.  These characters divide the square into contiguous regions.
  (Note that backslash characters are escaped, so a \ is represented as "\\".)
  Return the number of regions.

  Example 1:
    Input:
    [
      " /",
      "/ "
    ]
    Output: 2
    Explanation: The 2x2 grid is as follows:

  Example 2:
    Input:
    [
      " /",
      "  "
    ]
    Output: 1
    Explanation: The 2x2 grid is as follows:

  Example 3:
    Input:
    [
      "\\/",
      "/\\"
    ]
    Output: 4
    Explanation: (Recall that because \ characters are escaped, "\\/" refers to \/, and "/\\" refers to /\.)
    The 2x2 grid is as follows:

  Example 4:
    Input:
    [
      "/\\",
      "\\/"
    ]
    Output: 5
    Explanation: (Recall that because \ characters are escaped, "/\\" refers to /\, and "\\/" refers to \/.)
    The 2x2 grid is as follows:

  Example 5:
    Input:
    [
      "//",
      "/ "
    ]
    Output: 3
    Explanation: The 2x2 grid is as follows:

  Note:
    1. 1 <= grid.length == grid[0].length <= 30
    2. grid[i][j] is either '/', '\', or ' '.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/regions-cut-by-slashes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func regionsBySlashes(grid []string) int {
	n := len(grid)
	size := 4 * n * n
	uf := newUnionFind(size)
	for i := range grid {
		for j := range grid[i] {
			start := 4 * (n*j + i)
			up, right, down, left := start, start+1, start+2, start+3
			if grid[i][j] == '/' {
				uf.union(right, down)
				uf.union(up, left)
			} else if grid[i][j] == '\\' {
				uf.union(up, right)
				uf.union(down, left)
			} else {
				uf.union(right, up)
				uf.union(down, up)
				uf.union(left, up)
			}
			if i+1 < n {
				rightStart := start + 4
				rightLeft := rightStart + 3
				uf.union(rightLeft, right)
			}
			if j+1 < n {
				downStart := start + 4*n
				downUp := downStart
				uf.union(downUp, down)
			}
		}
	}
	out := 0
	mark := make([]bool, size)
	for i := 0; i < size; i++ {
		x := uf.find(i)
		if !mark[x] {
			out++
			mark[x] = true
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
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}
