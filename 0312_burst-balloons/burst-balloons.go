package main

/*
  有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
  现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
  求所能获得硬币的最大数量。

  说明:
    你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
    0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/burst-balloons
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func maxCoins(nums []int) int {
	n := len(nums)
	vals := make([]int, n+2)
	copy(vals[1:n+1], nums)
	vals[0], vals[n+1] = 1, 1
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for l := n - 1; l >= 0; l-- {
		for r := l + 2; r < n+2; r++ {
			for m := l + 1; m < r; m++ {
				tmp := vals[l]*vals[m]*vals[r] + dp[l][m] + dp[m][r]
				dp[l][r] = Max(dp[l][r], tmp)
			}
		}
	}
	return dp[0][n+1]
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
