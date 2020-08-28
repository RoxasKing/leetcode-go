package main

/*
  一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
  机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
  现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	dp := make([]int, len(obstacleGrid[0]))
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for i := range obstacleGrid {
		for j := range obstacleGrid[0] {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 && obstacleGrid[i][j-1] == 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[len(dp)-1]
}
