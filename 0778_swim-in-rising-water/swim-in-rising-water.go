package main

import "container/heap"

/*
  在一个 N x N 的坐标方格 grid 中，每一个方格的值 grid[i][j] 表示在位置 (i,j) 的平台高度。

  现在开始下雨了。当时间为 t 时，此时雨水导致水池中任意位置的水位为 t 。你可以从一个平台游向四周相邻的任意一个平台，但是前提是此时水位必须同时淹没这两个平台。假定你可以瞬间移动无限距离，也就是默认在方格内部游动是不耗时的。当然，在你游泳的时候你必须待在坐标方格里面。

  你从坐标方格的左上平台 (0，0) 出发。最少耗时多久你才能到达坐标方格的右下平台 (N-1, N-1)？

  提示:
    2 <= N <= 50.
    grid[i][j] 位于区间 [0, ..., N*N - 1] 内。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/swim-in-rising-water
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search + DFS
func swimInWater(grid [][]int) int {
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

var moves = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// Priority Queue(Heap Sort) + BFS
func swimInWater2(grid [][]int) int {
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

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
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

// Union-Find + Binary Search
func swimInWater3(grid [][]int) int {
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
			for _, m := range moves {
				row, col := i+m[0], j+m[1]
				if 0 <= row && row < n && 0 <= col && col < n && grid[row][col] <= limit {
					p1, p2 := uf.find(i*n+j), uf.find(row*n+col)
					if p1 != p2 {
						uf.union(p1, p2)
					}
				}
			}
		}
	}
	return uf.find(0) == uf.find(n*n-1)
}

type unionFind struct {
	ancestor []int
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := 0; i < n; i++ {
		ancestor[i] = i
	}
	return unionFind{ancestor}
}

func (uf unionFind) find(x int) int {
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(from, to int) {
	uf.ancestor[uf.find(from)] = uf.find(to)
}
