package My_LeetCode_In_Go

/*
  一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
  机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
  现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	cur := make([]int, len(obstacleGrid[0]))
	for i := range cur {
		if obstacleGrid[0][i] == 1 {
			break
		}
		cur[i] = 1
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := range obstacleGrid[i] {
			if obstacleGrid[i][j] == 1 {
				cur[j] = 0
			} else if j > 0 {
				cur[j] += cur[j-1]
			}
		}
	}
	return cur[len(cur)-1]
}
