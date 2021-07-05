package main

// Tags:
// DFS
func findCircleNum(M [][]int) int {
	n := len(M)
	visited := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(M, n, visited, i)
			count++
		}
	}
	return count
}

func dfs(M [][]int, N int, visited []bool, i int) {
	for j := 0; j < N; j++ {
		if M[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(M, N, visited, j)
		}
	}
}

// Union-Find
func findCircleNum2(M [][]int) int {
	n := len(M)
	uf := newUnionFind(n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	mark := make([]bool, len(M))
	count := 0
	for i := 0; i < n; i++ {
		index := uf.find(i)
		if !mark[index] {
			count++
			mark[index] = true
		}
	}
	return count
}

type unionFind struct {
	ancestor []int
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := range ancestor {
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
