package main

import "container/heap"

/*
  On an N x N grid, each square grid[i][j] represents the elevation at that point (i,j).
  Now rain starts to fall. At time t, the depth of the water everywhere is t. You can swim from a square to another 4-directionally adjacent square if and only if the elevation of both squares individually are at most t. You can swim infinite distance in zero time. Of course, you must stay within the boundaries of the grid during your swim.
  You start at the top left square (0, 0). What is the least time until you can reach the bottom right square (N-1, N-1)?

  Example 1:
    Input: [[0,2],[1,3]]
    Output: 3
    Explanation:
    At time 0, you are in grid location (0, 0).
    You cannot go anywhere else because 4-directionally adjacent neighbors have a higher elevation than t = 0.
    You cannot reach point (1, 1) until time 3.
    When the depth of water is 3, we can swim anywhere inside the grid.

  Example 2:
    Input: [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
    Output: 16
    Explanation:
     0  1  2  3  4
    24 23 22 21  5
    12 13 14 15 16
    11 17 18 19 20
    10  9  8  7  6
    The final route is marked in bold.
    We need to wait until time 16 so that (0, 0) and (4, 4) are connected.

  Note:
    2 <= N <= 50.
    grid[i][j] is a permutation of [0, ..., N*N - 1].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/swim-in-rising-water
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find + Binary Search
func swimInWater(grid [][]int) int {
	n := len(grid)
	l, r := Max(grid[0][0], grid[n-1][n-1]), n*n-1
	for l < r {
		m := l + (r-l)>>1
		if !checkByUnionFind(grid, m) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func checkByUnionFind(grid [][]int, limit int) bool {
	n := len(grid)
	uf := newUnionFind(n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > limit {
				continue
			}
			for _, move := range moves {
				x, y := i+move[0], j+move[1]
				if x < 0 || n-1 < x || y < 0 || n-1 < y || grid[x][y] > limit {
					continue
				}
				a := uf.find(i*n + j)
				b := uf.find(x*n + y)
				if a == b {
					continue
				}
				uf.union(a, b)
			}
		}
	}
	return uf.find(0) == uf.find(n*n-1)
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

var moves = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Binary Search + DFS
func swimInWater2(grid [][]int) int {
	n := len(grid)
	l, r := grid[0][0], n*n-1
	for l < r {
		m := l + (r-l)>>1
		if !checkByDFS(grid, m) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func checkByDFS(grid [][]int, limit int) bool {
	visited := [50][50]bool{}
	var dfs func(int, int) bool
	dfs = func(row, col int) bool {
		visited[row][col] = true
		for _, m := range moves {
			newRow, newCol := row+m[0], col+m[1]
			if 0 <= newRow && newRow < len(grid) && 0 <= newCol && newCol < len(grid[0]) &&
				!visited[newRow][newCol] && grid[newRow][newCol] <= limit {
				if newRow == len(grid)-1 && newCol == len(grid[0])-1 || dfs(newRow, newCol) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, 0)
}

// Priority Queue(Heap Sort) + BFS
func swimInWater3(grid [][]int) int {
	n := len(grid)
	visited := [50][50]bool{}
	var out int
	ph := &posHeap{}
	heap.Push(ph, pos{height: grid[0][0], row: 0, col: 0})
	for ph.Len() != 0 {
		cur := heap.Pop(ph).(pos)
		out = Max(out, grid[cur.row][cur.col])
		if cur.row == n-1 && cur.col == n-1 {
			break
		}
		for _, m := range moves {
			row, col := cur.row+m[0], cur.col+m[1]
			if 0 <= row && row < n && 0 <= col && col < n && !visited[row][col] {
				heap.Push(ph, pos{height: grid[row][col], row: row, col: col})
				visited[row][col] = true
			}
		}
	}
	return out
}

type pos struct {
	height int
	row    int
	col    int
}

type posHeap []pos

func (p posHeap) Len() int            { return len(p) }
func (p posHeap) Less(i, j int) bool  { return p[i].height < p[j].height }
func (p posHeap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *posHeap) Push(x interface{}) { *p = append(*p, x.(pos)) }
func (p *posHeap) Pop() interface{} {
	last := len(*p) - 1
	pos := (*p)[last]
	*p = (*p)[:last]
	return pos
}
