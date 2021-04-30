package main

/*
  Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

  Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

  Example 1:
    [[0,0,1,0,0,0,0,1,0,0,0,0,0],
     [0,0,0,0,0,0,0,1,1,1,0,0,0],
     [0,1,1,0,1,0,0,0,0,0,0,0,0],
     [0,1,0,0,1,1,0,0,1,0,1,0,0],
     [0,1,0,0,1,1,0,0,1,1,1,0,0],
     [0,0,0,0,0,0,0,0,0,0,1,0,0],
     [0,0,0,0,0,0,0,1,1,1,0,0,0],
     [0,0,0,0,0,0,0,1,1,0,0,0,0]]
    Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.

  Example 2:
    [[0,0,0,0,0,0,0,0]]
    Given the above grid, return 0.

  Note: The length of each dimension in the given grid does not exceed 50.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/max-area-of-island
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	uf := NewUnionFind(m * n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			x := i*n + j
			if i+1 < m && grid[i+1][j] == 1 {
				y := (i+1)*n + j
				uf.Union(x, y)
			}
			if j+1 < n && grid[i][j+1] == 1 {
				y := i*n + j + 1
				uf.Union(x, y)
			}
		}
	}

	out := 0
	cnt := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		r, c := i/n, i%n
		if grid[r][c] == 0 {
			continue
		}
		x := uf.Find(i)
		cnt[x]++
		if cnt[x] > out {
			out = cnt[x]
		}
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
