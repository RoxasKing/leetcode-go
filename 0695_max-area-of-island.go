package main

/*
  给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
  找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
*/

// DFS
func maxAreaOfIsland(grid [][]int) int {
	steps := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(int, int) int
	dfs = func(row, col int) int {
		if row < 0 || row == len(grid) || col < 0 || col == len(grid[0]) || grid[row][col] == 0 {
			return 0
		}
		grid[row][col] = 0
		res := 1
		for _, step := range steps {
			res += dfs(row+step[0], col+step[1])
		}
		return res
	}
	var max int
	for i := range grid {
		for j := range grid[0] {
			max = Max(max, dfs(i, j))
		}
	}
	return max
}

// DFS + Stack
func maxAreaOfIsland2(grid [][]int) int {
	steps := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var stack [][]int
	var elem []int
	var max, cur int
	for i := range grid {
		for j := range grid[0] {
			stack = append(stack, []int{i, j})
			for len(stack) != 0 {
				elem, stack = stack[len(stack)-1], stack[:len(stack)-1]
				row, col := elem[0], elem[1]
				if row < 0 || row == len(grid) ||
					col < 0 || col == len(grid[0]) ||
					grid[row][col] == 0 {
					continue
				}
				cur++
				grid[row][col] = 0
				for _, step := range steps {
					stack = append(stack, []int{row + step[0], col + step[1]})
				}
			}
			max, cur = Max(max, cur), 0
		}
	}
	return max
}

// BFS
func maxAreaOfIsland3(grid [][]int) int {
	steps := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var queue [][]int
	var elem []int
	var max, cur int
	for i := range grid {
		for j := range grid[0] {
			queue = append(queue, []int{i, j})
			for len(queue) != 0 {
				elem, queue = queue[0], queue[1:]
				row, col := elem[0], elem[1]
				if row < 0 || row == len(grid) ||
					col < 0 || col == len(grid[0]) ||
					grid[row][col] == 0 {
					continue
				}
				cur++
				grid[row][col] = 0
				for _, step := range steps {
					queue = append(queue, []int{row + step[0], col + step[1]})
				}
			}
			max, cur = Max(max, cur), 0
		}
	}
	return max
}
