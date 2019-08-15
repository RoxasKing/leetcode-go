package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	grid := make([][]int, len(obstacleGrid))
	for i := range grid {
		grid[i] = make([]int, len(obstacleGrid[0]))
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			// 如果有障碍物，则路径为 0
			if obstacleGrid[i][j] == 1 {
				grid[i][j] = 0
				continue
			}
			if i == 0 || j == 0 {
				if i > 0 && grid[i-1][j] == 0 || j > 0 && grid[i][j-1] == 0 {
					// 上方格子 或 左边格子 路径为 0 ，则当前格子也为 0
					grid[i][j] = 0
				} else {
					// 到达边界的路径数都 1
					grid[i][j] = 1
				}
				continue
			}
			// 到达 (i,j) 格子的路径数目，等于到达 上方格子 和 左边格子 路径数之和
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func main() {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{1, 0, 0},
	}
	//obstacleGrid := [][]int{
	//	{1, 0},
	//}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
