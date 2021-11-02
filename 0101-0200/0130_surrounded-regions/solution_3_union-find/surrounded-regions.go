package main

// Difficulty:
// Medium

// Tags:
// Union-Find

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	dummy := m * n
	uf := NewUnionFind(dummy + 1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != 'O' {
				continue
			}
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				uf.Union(i*n+j, dummy)
			}
			if i+1 < m && board[i+1][j] == 'O' {
				uf.Union(i*n+j, (i+1)*n+j)
			}
			if j+1 < n && board[i][j+1] == 'O' {
				uf.Union(i*n+j, i*n+j+1)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && uf.Find(i*n+j) != dummy {
				board[i][j] = 'X'
			}
		}
	}
}

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return UnionFind{parent: parent}
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
	if x > y {
		x, y = y, x
	}
	uf.parent[x] = y
}
