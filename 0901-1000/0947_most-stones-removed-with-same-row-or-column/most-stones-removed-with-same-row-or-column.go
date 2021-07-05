package main

// Tags:
// Union-Find
func removeStones(stones [][]int) int {
	stonesInRow := make(map[int][]int)
	stonesInCol := make(map[int][]int)
	for i, stone := range stones {
		stonesInRow[stone[0]] = append(stonesInRow[stone[0]], i)
		stonesInCol[stone[1]] = append(stonesInCol[stone[1]], i)
	}
	n := len(stones)
	uf := newUnionFind(n)
	for _, sts := range stonesInRow {
		for i := 1; i < len(sts); i++ {
			uf.union(sts[i-1], sts[i])
		}
	}
	for _, sts := range stonesInCol {
		for i := 1; i < len(sts); i++ {
			uf.union(sts[i-1], sts[i])
		}
	}
	mark := make(map[int]bool)
	count := 0
	for i := range stones {
		ancestorI := uf.find(i)
		if !mark[ancestorI] {
			count++
			mark[ancestorI] = true
		}
	}
	return n - count
}

type unionFind struct {
	ancestor []int
	isEnd    []bool
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := range ancestor {
		ancestor[i] = i
	}
	isEnd := make([]bool, n)
	return unionFind{ancestor: ancestor, isEnd: isEnd}
}

func (uf unionFind) find(x int) int {
	if uf.isEnd[uf.ancestor[x]] {
		return uf.ancestor[x]
	}
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
		uf.isEnd[x] = false
		uf.isEnd[uf.ancestor[x]] = true
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(x, y int) {
	ancestorX := uf.find(x)
	ancestorY := uf.find(y)
	uf.ancestor[ancestorX] = ancestorY
	uf.isEnd[ancestorX] = false
	uf.isEnd[ancestorY] = true
}

// DFS
func removeStones2(stones [][]int) int {
	stonesInRow := make(map[int][]int)
	stonesInCol := make(map[int][]int)
	for i, stone := range stones {
		stonesInRow[stone[0]] = append(stonesInRow[stone[0]], i)
		stonesInCol[stone[1]] = append(stonesInCol[stone[1]], i)
	}
	visitedRow := make(map[int]bool)
	visitedCol := make(map[int]bool)
	n := len(stones)
	visited := make([]bool, n)
	count := 0
	for i := n - 1; i >= 0; i-- {
		if visitedRow[stones[i][0]] || visitedCol[stones[i][1]] || visited[i] {
			continue
		}
		count++
		dfs(stones, stonesInRow, stonesInCol, visitedRow, visitedCol, visited, i)
	}
	return n - count
}

func dfs(
	stones [][]int,
	stonesInRow, stonesInCol map[int][]int,
	visitedRow, visitedCol map[int]bool,
	visited []bool, i int,
) {
	if !visitedRow[stones[i][0]] {
		visitedRow[stones[i][0]] = true
		for _, stoneIdx := range stonesInRow[stones[i][0]] {
			if !visited[stoneIdx] {
				visited[stoneIdx] = true
				dfs(stones, stonesInRow, stonesInCol, visitedRow, visitedCol, visited, stoneIdx)
			}
		}
	}
	if !visitedCol[stones[i][1]] {
		visitedCol[stones[i][1]] = true
		for _, stoneIdx := range stonesInCol[stones[i][1]] {
			if !visited[stoneIdx] {
				visited[stoneIdx] = true
				dfs(stones, stonesInRow, stonesInCol, visitedRow, visitedCol, visited, stoneIdx)
			}
		}
	}
}
