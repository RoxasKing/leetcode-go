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

func maxCoins(nums []int) int {
	val := make([]int, len(nums)+2)
	for i := 1; i <= len(nums); i++ {
		val[i] = nums[i-1]
	}
	val[0], val[len(val)-1] = 1, 1
	rec := make([][]int, len(nums)+2)
	for i := range rec {
		rec[i] = make([]int, len(nums)+2)
		for j := 0; j < len(rec[i]); j++ {
			rec[i][j] = -1
		}
	}
	var solve func(int, int) int
	solve = func(l, r int) int {
		if l+1 >= r {
			return 0
		}
		if rec[l][r] != -1 {
			return rec[l][r]
		}
		for i := l + 1; i < r; i++ {
			sum := val[l]*val[i]*val[r] + solve(l, i) + solve(i, r)
			rec[l][r] = Max(rec[l][r], sum)
		}
		return rec[l][r]
	}
	return solve(0, len(nums)+1)
}

// Dynamic Programming
func maxCoins2(nums []int) int {
	val := make([]int, len(nums)+2)
	copy(val[1:len(nums)+1], nums)
	val[0], val[len(nums)+1] = 1, 1
	rec := make([][]int, len(nums)+2)
	for i := range rec {
		rec[i] = make([]int, len(nums)+2)
	}
	for l := len(nums) - 1; l >= 0; l-- {
		for r := l + 2; r <= len(nums)+1; r++ {
			for m := l + 1; m < r; m++ {
				sum := val[l]*val[m]*val[r] + rec[l][m] + rec[m][r]
				rec[l][r] = Max(rec[l][r], sum)
			}
		}
	}
	return rec[0][len(nums)+1]
}
