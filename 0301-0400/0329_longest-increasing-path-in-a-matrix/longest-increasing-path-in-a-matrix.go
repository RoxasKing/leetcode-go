package main

/*
  Given an m x n integers matrix, return the length of the longest increasing path in matrix.

  From each cell, you can either move in four directions: left, right, up, or down. You may not move diagonally or move outside the boundary (i.e., wrap-around is not allowed).

  Example 1:
    Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
    Output: 4
    Explanation: The longest increasing path is [1, 2, 6, 9].

  Example 2:
    Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
    Output: 4
    Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.

  Example 3:
    Input: matrix = [[1]]
    Output: 1

  Constraints:
    1. m == matrix.length
    2. n == matrix[i].length
    3. 1 <= m, n <= 200
    4. 0 <= matrix[i][j] <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Topological Sorting
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	size := m * n
	edges := make([][]int, size)
	indeg := make([]int, size)
	for i := 0; i < size; i++ {
		r, c := i/n, i%n
		for _, f := range forwards {
			nr, nc := r+f[0], c+f[1]
			if nr < 0 || m-1 < nr || nc < 0 || n-1 < nc || matrix[nr][nc] <= matrix[r][c] {
				continue
			}
			j := nr*n + nc
			edges[i] = append(edges[i], j)
			indeg[j]++
		}
	}

	q := [][2]int{}
	for i := 0; i < size; i++ {
		if indeg[i] == 0 {
			q = append(q, [2]int{i, 1})
		}
	}

	out := 0
	for len(q) > 0 {
		u, dist := q[0][0], q[0][1]
		q = q[1:]
		out = Max(out, dist)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, [2]int{v, dist + 1})
			}
		}
	}
	return out
}

var forwards = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
