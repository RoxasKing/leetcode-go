package main

/*
  为了给刷题的同学一些奖励，力扣团队引入了一个弹簧游戏机。游戏机由 N 个特殊弹簧排成一排，编号为 0 到 N-1。初始有一个小球在编号 0 的弹簧处。若小球在编号为 i 的弹簧处，通过按动弹簧，可以选择把小球向右弹射 jump[i] 的距离，或者向左弹射到任意左侧弹簧的位置。也就是说，在编号为 i 弹簧处按动弹簧，小球可以弹向 0 到 i-1 中任意弹簧或者 i+jump[i] 的弹簧（若 i+jump[i]>=N ，则表示小球弹出了机器）。小球位于编号 0 处的弹簧时不能再向左弹。

  为了获得奖励，你需要将小球弹出机器。请求出最少需要按动多少次弹簧，可以将小球从编号 0 弹簧弹出整个机器，即向右越过编号 N-1 的弹簧。

  示例 1：
    输入：jump = [2, 5, 1, 1, 1, 1]
    输出：3
    解释：小 Z 最少需要按动 3 次弹簧，小球依次到达的顺序为 0 -> 2 -> 1 -> 6，最终小球弹出了机器。

  限制：
    1. 1 <= jump.length <= 10^6
    2. 1 <= jump[i] <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/zui-xiao-tiao-yue-ci-shu
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func minJump(jump []int) int {
	n := len(jump)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = n
	}
	dp[0] = 0

	out := 1<<31 - 1
	for i := range dp {
		j := i + jump[i]
		if j >= n {
			out = Min(out, dp[i]+1)
			continue
		}
		if dp[j] <= dp[i]+1 {
			continue
		}
		dp[j] = dp[i] + 1
		for k := j - 1; dp[k] >= dp[j]+1; k-- {
			dp[k] = dp[j] + 1
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Dynamic Programming
func minJump2(jump []int) int {
	n := len(jump)
	maxDist := make([]int, n)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = n
	}

	out := n
	times := 0
	for i := 0; i < n; i++ {
		if i > maxDist[times] {
			times++
		}
		dp[i] = Min(dp[i], times+1)

		next := i + jump[i]
		if next >= n {
			out = Min(out, dp[i]+1)
		} else {
			maxDist[dp[i]+1] = Max(maxDist[dp[i]+1], next)
			dp[next] = Min(dp[next], dp[i]+1)
		}
	}
	return out
}
