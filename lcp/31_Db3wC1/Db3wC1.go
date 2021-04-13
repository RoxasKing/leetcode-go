package main

/*
  某解密游戏中，有一个 N*M 的迷宫，迷宫地形会随时间变化而改变，迷宫出口一直位于 (n-1,m-1) 位置。迷宫变化规律记录于 maze 中，maze[i] 表示 i 时刻迷宫的地形状态，"." 表示可通行空地，"#" 表示陷阱。

  地形图初始状态记作 maze[0]，此时小力位于起点 (0,0)。此后每一时刻可选择往上、下、左、右其一方向走一步，或者停留在原地。

  小力背包有以下两个魔法卷轴（卷轴使用一次后消失）：
    1. 临时消除术：将指定位置在下一个时刻变为空地；
    2. 永久消除术：将指定位置永久变为空地。
  请判断在迷宫变化结束前（含最后时刻），小力能否在不经过任意陷阱的情况下到达迷宫出口呢？

  注意： 输入数据保证起点和终点在所有时刻均为空地。

  示例 1：
    输入：maze = [[".#.","#.."],["...",".#."],[".##",".#."],["..#",".#."]]
    输出：true
    解释：

  示例 2：
    输入：maze = [[".#.","..."],["...","..."]]
    输出：false
    解释：由于时间不够，小力无法到达终点逃出迷宫。

  示例 3：
    输入：maze = [["...","...","..."],[".##","###","##."],[".##","###","##."],[".##","###","##."],[".##","###","##."],[".##","###","##."],[".##","###","##."]]
    输出：false
    解释：由于道路不通，小力无法到达终点逃出迷宫。

  提示：
    1. 1 <= maze.length <= 100
    2. 1 <= maze[i].length, maze[i][j].length <= 50
    3. maze[i][j] 仅包含 "."、"#"

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/Db3wC1
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming + DFS
func escapeMaze(maze [][]string) bool {
	t, m, n := len(maze), len(maze[0]), len(maze[0][0])
	dp := make([][][][2][2][2]bool, t) // [times][row][col][scroll1][scroll2][isStop]
	for i := range dp {
		dp[i] = make([][][2][2][2]bool, m)
		for j := range dp[i] {
			dp[i][j] = make([][2][2][2]bool, n)
		}
	}

	dfs(maze, t, m, n, dp, 0, 0, 0, 0, 0, 0)

	for i := 0; i < t; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				for l := 0; l < 2; l++ {
					if dp[i][m-1][n-1][j][k][l] {
						return true
					}
				}
			}
		}
	}
	return false
}

func dfs(maze [][]string, t, m, n int, dp [][][][2][2][2]bool, times, row, col int, scroll1, scroll2, isStop int) {
	if dp[times][row][col][scroll1][scroll2][isStop] {
		return
	}

	dp[times][row][col][scroll1][scroll2][isStop] = true

	if times == t-1 {
		return
	}

	if isStop == 1 || maze[times+1][row][col] == '.' {
		dfs(maze, t, m, n, dp, times+1, row, col, scroll1, scroll2, isStop)
	}

	for _, f := range forwards {
		i, j := row+f[0], col+f[1]
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			continue
		}
		if maze[times+1][i][j] == '.' {
			dfs(maze, t, m, n, dp, times+1, i, j, scroll1, scroll2, 0)
			continue
		}
		if scroll1 == 0 {
			dfs(maze, t, m, n, dp, times+1, i, j, 1, scroll2, 0)
		}
		if scroll2 == 0 {
			dfs(maze, t, m, n, dp, times+1, i, j, scroll1, 1, 1)
		}
	}
}

var forwards = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
