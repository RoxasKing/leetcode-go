package leetcode

/*
  在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。
  每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单元格 (i, j) 上。
  请你返回最终形体的表面积。
*/

func surfaceArea(grid [][]int) int {
	steps := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var area int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] != 0 {
				area += 2
			}
			for _, step := range steps {
				r, c := i+step[0], j+step[1]
				if r < 0 || r > len(grid)-1 || c < 0 || c > len(grid)-1 {
					area += grid[i][j]
				} else if grid[r][c] < grid[i][j] {
					area += grid[i][j] - grid[r][c]
				}
			}
		}
	}
	return area
}
