package main

/*
  有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。
  省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
  给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。
  返回矩阵中 省份 的数量。

  示例 1：
    输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
    输出：2

  示例 2：
    输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
    输出：3

  提示：
    1 <= n <= 200
    n == isConnected.length
    n == isConnected[i].length
    isConnected[i][j] 为 1 或 0
    isConnected[i][i] == 1
    isConnected[i][j] == isConnected[j][i]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-provinces
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
