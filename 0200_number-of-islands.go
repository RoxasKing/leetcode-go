package leetcode

/*
  给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
  岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
  此外，你可以假设该网格的四条边均被水包围。
*/

// BFS + Stack
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	steps := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var count int
	var queue [][]int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '0' {
				continue
			}
			count++
			grid[i][j] = '0'
			queue = append(queue, []int{i, j})
			for len(queue) != 0 {
				i, j := queue[0][0], queue[0][1]
				for _, step := range steps {
					ni, nj := i+step[0], j+step[1]
					if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[0]) ||
						grid[ni][nj] == '0' {
						continue
					}
					grid[ni][nj] = '0'
					queue = append(queue, []int{ni, nj})
				}
				queue = queue[1:]
			}
		}
	}
	return count
}

// DFS + Recursion
func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	steps := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(int, int)
	dfs = func(i, j int) {
		for _, step := range steps {
			ni, nj := i+step[0], j+step[1]
			if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[0]) ||
				grid[ni][nj] == '0' {
				continue
			}
			grid[ni][nj] = '0'
			dfs(ni, nj)
		}
	}
	var count int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '0' {
				continue
			}
			count++
			grid[i][j] = '0'
			dfs(i, j)
		}
	}
	return count
}
