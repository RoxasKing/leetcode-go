package main

/*
  给定一个表示分数的非负整数数组。 玩家 1 从数组任意一端拿取一个分数，随后玩家 2 继续从剩余数组任意一端拿取分数，然后玩家 1 拿，…… 。每次一个玩家只能拿取一个分数，分数被拿取之后不再可取。直到没有剩余分数可取时游戏结束。最终获得分数总和最多的玩家获胜。

  给定一个表示分数的数组，预测玩家1是否会成为赢家。你可以假设每个玩家的玩法都会使他的分数最大化。

  提示：
    1 <= 给定的数组长度 <= 20.
    数组里所有分数都为非负数且不会大于 10000000 。
    如果最终两个玩家的分数相等，那么玩家 1 仍为赢家。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/predict-the-winner
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Recursion
func PredictTheWinner(nums []int) bool {
	return maxScore(nums, 1) >= 0
}

func maxScore(nums []int, turn int) int {
	if len(nums) == 1 {
		return nums[0] * turn
	}
	head, tail := 0, len(nums)-1
	pickHead := nums[head]*turn + maxScore(nums[head+1:], -turn)
	pickTail := nums[tail]*turn + maxScore(nums[:tail], -turn)
	return Max(pickHead*turn, pickTail*turn) * turn
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Dynamic Programming
func PredictTheWinner2(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = nums[i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[j] = Max(nums[i]-dp[j], nums[j]-dp[j-1])
		}
	}
	return dp[n-1] >= 0
}
