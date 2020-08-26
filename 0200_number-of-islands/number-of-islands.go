package main

/*
  给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
  岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
  此外，你可以假设该网格的四条边均被水包围。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-islands
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS + Stack
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var count int
	var queue [][]int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] != '1' {
				continue
			}
			count++
			grid[i][j] = '#'
			queue = append(queue, []int{i, j})
			for len(queue) != 0 {
				row, col := queue[0][0], queue[0][1]
				queue = queue[1:]
				for _, move := range moves {
					newRow, newCol := row+move[0], col+move[1]
					if newRow < 0 || newRow >= len(grid) ||
						newCol < 0 || newCol >= len(grid[0]) ||
						grid[newRow][newCol] != '1' {
						continue
					}
					grid[newRow][newCol] = '#'
					queue = append(queue, []int{newRow, newCol})
				}
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
	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(int, int)
	dfs = func(row, col int) {
		for _, move := range moves {
			newRow, newCol := row+move[0], col+move[1]
			if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) ||
				grid[newRow][newCol] != '1' {
				continue
			}
			grid[newRow][newCol] = '#'
			dfs(newRow, newCol)
		}
	}
	var count int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] != '1' {
				continue
			}
			count++
			grid[i][j] = '#'
			dfs(i, j)
		}
	}
	return count
}
