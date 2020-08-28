package main

/*
  给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
*/

func maxProduct(nums []int) int {
	curMax, curMin := nums[0], nums[0]
	max := curMax
	for i := 1; i < len(nums); i++ {
		a := curMax * nums[i]
		b := curMin * nums[i]
		curMax = Max(Max(a, b), nums[i])
		curMin = Min(Min(a, b), nums[i])
		max = Max(max, curMax)
	}
	return max
}
