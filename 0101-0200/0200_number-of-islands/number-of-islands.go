package main

/*
  Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

  An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

  Example 1:
    Input: grid = [
      ["1","1","1","1","0"],
      ["1","1","0","1","0"],
      ["1","1","0","0","0"],
      ["0","0","0","0","0"]
    ]
    Output: 1

  Example 2:
    Input: grid = [
      ["1","1","0","0","0"],
      ["1","1","0","0","0"],
      ["0","0","1","0","0"],
      ["0","0","0","1","1"]
    ]
    Output: 3

  Constraints:
    1. m == grid.length
    2. n == grid[i].length
    3. 1 <= m, n <= 300
    4. grid[i][j] is '0' or '1'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-islands
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	uf := NewUnionFind(m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			x := i*n + j
			if i < m-1 && grid[i+1][j] == '1' {
				y := (i+1)*n + j
				uf.Union(x, y)
			}
			if j < n-1 && grid[i][j+1] == '1' {
				y := i*n + j + 1
				uf.Union(x, y)
			}
		}
	}
	out := 0
	visited := make([]bool, m*n)
	for i := 0; i < m*n; i++ {
		r, c := i/n, i%n
		if grid[r][c] == '0' {
			continue
		}
		x := uf.Find(i)
		if visited[x] {
			continue
		}
		out++
		visited[x] = true
	}
	return out
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return UnionFind{
		parent: parent,
		size:   size,
	}
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
	uf.size[x] += uf.size[y]
}
