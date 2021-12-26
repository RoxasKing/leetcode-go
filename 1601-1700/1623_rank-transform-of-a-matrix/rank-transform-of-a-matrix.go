package main

import "sort"

// Tags:
// Union-Find
// Topological Sort
// BFS

func matrixRankTransform(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	idxC := make([]int, n)
	for i := range idxC {
		idxC[i] = i
	}
	idxR := make([]int, m)
	for i := range idxR {
		idxR[i] = i
	}
	NewUnionFind(m * n)
	for k := 0; k < m; k++ {
		sort.Slice(idxC, func(i, j int) bool { return mat[k][idxC[i]] < mat[k][idxC[j]] })
		x := idxC[0]
		u := k*n + x
		for _, y := range idxC[1:] {
			v := k*n + y
			if mat[k][x] == mat[k][y] {
				Union(u, v)
			} else {
				x, u = y, v
			}
		}
	}
	for k := 0; k < n; k++ {
		sort.Slice(idxR, func(i, j int) bool { return mat[idxR[i]][k] < mat[idxR[j]][k] })
		x := idxR[0]
		u := x*n + k
		for _, y := range idxR[1:] {
			v := y*n + k
			if mat[x][k] == mat[y][k] {
				Union(u, v)
			} else {
				x, u = y, v
			}
		}
	}

	g := map[int][]int{}
	d := map[int]int{}
	for k := 0; k < m; k++ {
		sort.Slice(idxC, func(i, j int) bool { return mat[k][idxC[i]] < mat[k][idxC[j]] })
		x := idxC[0]
		for _, y := range idxC[1:] {
			if mat[k][x] == mat[k][y] {
				continue
			}
			u := Find(k*n + x)
			v := Find(k*n + y)
			g[u] = append(g[u], v)
			d[v]++
			x = y
		}
	}
	for k := 0; k < n; k++ {
		sort.Slice(idxR, func(i, j int) bool { return mat[idxR[i]][k] < mat[idxR[j]][k] })
		x := idxR[0]
		for _, y := range idxR[1:] {
			if mat[x][k] == mat[y][k] {
				continue
			}
			u := Find(x*n + k)
			v := Find(y*n + k)
			g[u] = append(g[u], v)
			d[v]++
			x = y
		}
	}

	added := map[int]bool{}
	q := [][]int{}
	for i := 0; i < m*n; i++ {
		x := Find(i)
		if added[x] || d[x] != 0 {
			continue
		}
		added[x] = true
		q = append(q, []int{x, 1})
	}

	memo := map[int]int{}
	for ; len(q) > 0; q = q[1:] {
		u, rank := q[0][0], q[0][1]
		memo[u] = rank
		for _, v := range g[u] {
			d[v]--
			if d[v] == 0 {
				q = append(q, []int{v, rank + 1})
			}
		}
	}

	out := make([][]int, m)
	for i := 0; i < m; i++ {
		out[i] = make([]int, n)
		for j := 0; j < n; j++ {
			x := Find(i*n + j)
			out[i][j] = memo[x]
		}
	}

	return out
}

var parent = [250000]int{}
var size = [250000]int{}

func NewUnionFind(n int) {
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
}

func Find(x int) int {
	if parent[x] != x {
		parent[x] = Find(parent[x])
	}
	return parent[x]
}

func Union(x, y int) {
	x = Find(x)
	y = Find(y)
	if x == y {
		return
	}
	if size[x] < size[y] {
		x, y = y, x
	}
	parent[y] = x
	size[x] += size[y]
}
