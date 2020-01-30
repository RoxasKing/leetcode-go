package My_LeetCode_In_Go

/*
  一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
  机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
  问总共有多少条不同的路径？
*/

func uniquePaths(m int, n int) int {
	cur := make([]int, n) // Save the number of paths for each layer
	for i := range cur {
		cur[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] += cur[j-1]
		}
	}
	return cur[n-1]
}
