package main

/*
  On a 2D plane, we place n stones at some integer coordinate points. Each coordinate point may have at most one stone.

  A stone can be removed if it shares either the same row or the same column as another stone that has not been removed.

  Given an array stones of length n where stones[i] = [xi, yi] represents the location of the ith stone, return the largest possible number of stones that can be removed.

  Example 1:
    Input: stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
    Output: 5
    Explanation: One way to remove 5 stones is as follows:
    1. Remove stone [2,2] because it shares the same row as [2,1].
    2. Remove stone [2,1] because it shares the same column as [0,1].
    3. Remove stone [1,2] because it shares the same row as [1,0].
    4. Remove stone [1,0] because it shares the same column as [0,0].
    5. Remove stone [0,1] because it shares the same row as [0,0].
    Stone [0,0] cannot be removed since it does not share a row/column with another stone still on the plane.

  Example 2:
    Input: stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
    Output: 3
    Explanation: One way to make 3 moves is as follows:
    1. Remove stone [2,2] because it shares the same row as [2,0].
    2. Remove stone [2,0] because it shares the same column as [0,0].
    3. Remove stone [0,2] because it shares the same row as [0,0].
    Stones [0,0] and [1,1] cannot be removed since they do not share a row/column with another stone still on the plane.

  Example 3:
    Input: stones = [[0,0]]
    Output: 0
    Explanation: [0,0] is the only stone on the plane, so you cannot remove it.

  Constraints:
    1 <= stones.length <= 1000
    0 <= xi, yi <= 104
    No two stones are at the same coordinate point.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/most-stones-removed-with-same-row-or-column
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
