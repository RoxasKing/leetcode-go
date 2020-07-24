package leetcode

/*
  爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。
  最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：
    选出任一 x，满足 0 < x < N 且 N % x == 0 。
    用 N - x 替换黑板上的数字 N 。
  如果玩家无法执行这些操作，就会输掉游戏。
  只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 false。假设两个玩家都以最佳状态参与游戏。

  提示：
    1 <= N <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/divisor-game
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func divisorGame(N int) bool {
	dp := make([]bool, N+5)
	dp[1], dp[2] = false, true
	for i := 3; i <= N; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[N]
}

func divisorGame2(N int) bool {
	return N%2 == 0
}
