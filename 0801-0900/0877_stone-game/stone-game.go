package main

/*
  亚历克斯和李用几堆石子在做游戏。偶数堆石子排成一行，每堆都有正整数颗石子 piles[i] 。

  游戏以谁手中的石子最多来决出胜负。石子的总数是奇数，所以没有平局。

  亚历克斯和李轮流进行，亚历克斯先开始。 每回合，玩家从行的开始或结束处取走整堆石头。 这种情况一直持续到没有更多的石子堆为止，此时手中石子最多的玩家获胜。

  假设亚历克斯和李都发挥出最佳水平，当亚历克斯赢得比赛时返回 true ，当李赢得比赛时返回 false 。

  提示：
    2 <= piles.length <= 500
    piles.length 是偶数。
    1 <= piles[i] <= 500
    sum(piles) 是奇数。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/stone-game
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func stoneGame(piles []int) bool {
	return true
}

// Dynamic Programming
func stoneGame2(piles []int) bool {
	n := len(piles)
	dp := make([]int, n)
	copy(dp, piles)
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i] = Max(piles[i]-dp[j], piles[j]-dp[j-1])
		}
	}
	return dp[n-1] >= 0
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
