package main

import "container/heap"

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

// Priority Queue + BFS
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
