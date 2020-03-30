package leetcode

/*
  你现在手里有一份大小为 N x N 的『地图』（网格） grid，上面的每个『区域』（单元格）都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地，你知道距离陆地区域最远的海洋区域是是哪一个吗？请返回该海洋区域到离它最近的陆地区域的距离。
  我们这里说的距离是『曼哈顿距离』（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个区域之间的距离是 |x0 - x1| + |y0 - y1| 。
  如果我们的地图上只有陆地或者海洋，请返回 -1。
*/

// Dynamic Programming
func maxDistance(grid [][]int) int {
	dp := make([][]int, len(grid)) // 储存当前海洋离陆地的最近距离
	for i := range dp {
		dp[i] = make([]int, len(grid))
		for j := range dp[i] {
			if grid[i][j] == 0 {
				dp[i][j] = 1<<31 - 1 // 海洋坐标初始化为int最大值
			} else {
				dp[i][j] = 0 // 陆地坐标初始化为0,代表从海洋到陆地的起点
			}
		}
	}

	for i := range grid {
		for j := range grid {
			if grid[i][j] == 1 {
				continue
			}
			if i-1 >= 0 {
				dp[i][j] = Min(dp[i][j], dp[i-1][j]+1)
			}
			if j-1 >= 0 {
				dp[i][j] = Min(dp[i][j], dp[i][j-1]+1)
			}
		}
	}

	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid) - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				continue
			}
			if i+1 < len(grid) {
				dp[i][j] = Min(dp[i][j], dp[i+1][j]+1)
			}
			if j+1 < len(grid) {
				dp[i][j] = Min(dp[i][j], dp[i][j+1]+1)
			}
		}
	}

	max := -1
	for i := range grid {
		for j := range grid {
			if grid[i][j] == 0 {
				max = Max(max, dp[i][j])
			}
		}
	}
	if max == 1<<31-1 {
		return -1
	}
	return max
}
