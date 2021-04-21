package main

/*
  给定一个整数矩阵，找出最长递增路径的长度。

  对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
	}
	var max int
	for i := range matrix {
		for j := range matrix[0] {
			if memo[i][j] == 0 {
				max = Max(max, dfs(matrix, memo, -1<<31, i, j))
			}
		}
	}
	return max
}

func dfs(matrix, memo [][]int, preVal, i, j int) int {
	if i < 0 || len(matrix)-1 < i || j < 0 || len(matrix[0])-1 < j ||
		matrix[i][j] <= preVal {
		return 0
	}
	if memo[i][j] > 0 {
		return memo[i][j]
	}
	memo[i][j]++
	memo[i][j] = Max(memo[i][j], dfs(matrix, memo, matrix[i][j], i-1, j)+1)
	memo[i][j] = Max(memo[i][j], dfs(matrix, memo, matrix[i][j], i+1, j)+1)
	memo[i][j] = Max(memo[i][j], dfs(matrix, memo, matrix[i][j], i, j-1)+1)
	memo[i][j] = Max(memo[i][j], dfs(matrix, memo, matrix[i][j], i, j+1)+1)
	return memo[i][j]
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Topological Sorting
func longestIncreasingPath2(matrix [][]int) int {
	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(matrix), len(matrix[0])
	indeg := make([][]int, m)
	for i := range indeg {
		indeg[i] = make([]int, n)
	}
	for i := range matrix {
		for j := range matrix[0] {
			for _, move := range moves {
				x, y := i+move[0], j+move[1]
				if 0 <= x && x < m && 0 <= y && y < n && matrix[x][y] < matrix[i][j] {
					indeg[i][j]++
				}
			}
		}
	}

	q := [][]int{}
	for i := range matrix {
		for j := range matrix[0] {
			if indeg[i][j] == 0 {
				q = append(q, []int{i, j})
			}
		}
	}

	cnt := 0
	for len(q) != 0 {
		cnt++
		size := len(q)
		for _, e := range q {
			i, j := e[0], e[1]
			for _, move := range moves {
				x, y := i+move[0], j+move[1]
				if 0 <= x && x < m && 0 <= y && y < n && matrix[x][y] > matrix[i][j] {
					indeg[x][y]--
					if indeg[x][y] == 0 {
						q = append(q, []int{x, y})
					}
				}
			}
		}
		q = q[size:]
	}
	return cnt
}

// Topological Sorting
func longestIncreasingPath3(matrix [][]int) int {
	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(matrix), len(matrix[0])
	size := m * n
	edges := make([][]int, size)
	indeg := make([]int, size)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k1 := i*n + j
			for _, mo := range moves {
				x, y := i+mo[0], j+mo[1]
				if 0 <= x && x < m && 0 <= y && y < n && matrix[x][y] < matrix[i][j] {
					k2 := x*n + y
					edges[k2] = append(edges[k2], k1)
					indeg[k1]++
				}
			}
		}
	}

	q := []int{}
	for i := 0; i < size; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	cnt := 0
	for len(q) > 0 {
		cnt++
		total := len(q)
		for _, a := range q {
			for _, b := range edges[a] {
				indeg[b]--
				if indeg[b] == 0 {
					q = append(q, b)
				}
			}
		}
		q = q[total:]
	}
	return cnt
}
