package main

/*
  班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。

  给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。

  提示：
    1 <= N <= 200
    M[i][i] == 1
    M[i][j] == M[j][i]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/friend-circles
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS
func findCircleNum(M [][]int) int {
	N := len(M)
	visited := make([]bool, N)
	var count int
	var dfs func(int)
	dfs = func(i int) {
		for j := 0; j < N; j++ {
			if M[i][j] == 1 && !visited[j] {
				visited[j] = true
				dfs(j)
			}
		}
	}
	for i := 0; i < N; i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}
	return count
}

// Union-Find
func findCircleNum2(M [][]int) int {
	uf := newUnionFind(len(M))
	for i := 0; i < len(M)-1; i++ {
		for j := i + 1; j < len(M); j++ {
			if M[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	var out int
	mark := make([]bool, len(M))
	for i := 0; i < len(M); i++ {
		index := uf.find(i)
		if !mark[index] {
			out++
			mark[index] = true
		}
	}
	return out
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
