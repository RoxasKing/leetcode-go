package main

/*
  The demons had captured the princess and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of m x n rooms laid out in a 2D grid. Our valiant knight was initially positioned in the top-left room and must fight his way through dungeon to rescue the princess.

  The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.

  Some of the rooms are guarded by demons (represented by negative integers), so the knight loses health upon entering these rooms; other rooms are either empty (represented as 0) or contain magic orbs that increase the knight's health (represented by positive integers).

  To reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.

  Return the knight's minimum initial health so that he can rescue the princess.

  Note that any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.

  Example 1:
    Input: dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
    Output: 7
    Explanation: The initial health of the knight must be at least 7 if he follows the optimal path: RIGHT-> RIGHT -> DOWN -> DOWN.

  Example 2:
    Input: dungeon = [[0]]
    Output: 1

  Constraints:
    1. m == dungeon.length
    2. n == dungeon[i].length
    3. 1 <= m, n <= 200
    4. -1000 <= dungeon[i][j] <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/dungeon-game
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([]int, n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			hp := 1<<31 - 1
			if i < m-1 {
				hp = Min(hp, dp[j])
			}
			if j < n-1 {
				hp = Min(hp, dp[j+1])
			}
			if hp == 1<<31-1 {
				hp = 1
			}
			dp[j] = hp
			if hp-dungeon[i][j] > 0 {
				dp[j] = hp - dungeon[i][j]
			}
		}
	}
	return dp[0]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
