package codinginterviews

/*
  在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

  提示：
    0 < grid.length <= 200
    0 < grid[0].length <= 200

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp := make([]int, len(grid[0]))
	dp[0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		for j := range grid[0] {
			if j != 0 {
				dp[j] = Max(dp[j-1], dp[j]) + grid[i][j]
			} else {
				dp[j] += grid[i][j]
			}
		}
	}
	return dp[len(dp)-1]
}
