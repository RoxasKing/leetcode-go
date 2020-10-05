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
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	outdegrees := make([][]int, len(matrix))
	for i := range outdegrees {
		outdegrees[i] = make([]int, len(matrix[0]))
	}
	for i := range matrix {
		for j := range matrix[0] {
			for _, move := range moves {
				newRow, newCol := i+move[0], j+move[1]
				if 0 <= newRow && newRow < len(matrix) &&
					0 <= newCol && newCol < len(matrix[0]) &&
					matrix[i][j] < matrix[newRow][newCol] {
					outdegrees[i][j]++
				}
			}
		}
	}
	var queue [][]int
	for i := range matrix {
		for j := range matrix[0] {
			if outdegrees[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	var count int
	for len(queue) != 0 {
		count++
		size := len(queue)
		for _, q := range queue {
			row, col := q[0], q[1]
			for _, move := range moves {
				newRow, newCol := row+move[0], col+move[1]
				if 0 <= newRow && newRow < len(matrix) &&
					0 <= newCol && newCol < len(matrix[0]) &&
					matrix[newRow][newCol] < matrix[row][col] {
					outdegrees[newRow][newCol]--
					if outdegrees[newRow][newCol] == 0 {
						queue = append(queue, []int{newRow, newCol})
					}
				}
			}
		}
		queue = queue[size:]
	}
	return count
}
