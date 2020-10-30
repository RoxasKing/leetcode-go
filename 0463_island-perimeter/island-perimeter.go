package main

/*
  给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。
  网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
  岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。

  示例 :
    输入:
    [[0,1,0,0],
     [1,1,1,0],
     [0,1,0,0],
     [1,1,0,0]]
    输出: 16
    解释: 它的周长是下面图片中的 16 个黄色的边：

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/island-perimeter
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func islandPerimeter(grid [][]int) int {
	var out int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				continue
			}
			for _, step := range steps {
				newI, newJ := i+step[0], j+step[1]
				if newI < 0 || len(grid)-1 < newI || newJ < 0 || len(grid[0])-1 < newJ || grid[newI][newJ] == 0 {
					out++
				}
			}
		}
	}
	return out
}

// DFS
func islandPerimeter2(grid [][]int) int {
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				var out int
				dfs(grid, i, j, &out)
				return out
			}
		}
	}
	return 0
}

func dfs(grid [][]int, row, col int, out *int) {
	grid[row][col] = -1
	for _, step := range steps {
		newR, newC := row+step[0], col+step[1]
		if newR < 0 || len(grid)-1 < newR || newC < 0 || len(grid[0])-1 < newC || grid[newR][newC] == 0 {
			*out++
			continue
		}
		if grid[newR][newC] == 1 {
			dfs(grid, newR, newC, out)
		}
	}
}

var steps = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
