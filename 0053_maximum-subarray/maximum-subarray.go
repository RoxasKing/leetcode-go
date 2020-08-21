package main

/*
  给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/

// Dynamic Programming
func maxSubArray(nums []int) int {
	cur, max := -1<<31, -1<<31
	for _, num := range nums {
		cur = Max(cur+num, num)
		max = Max(max, cur)
	}
	return max
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Dynamic Programming
func maxSubArrayReturnArray(nums []int) []int {
	cur, max := -1<<31, -1<<31
	var out []int
	var l, r int
	for r = range nums {
		cur += nums[r]
		if cur < nums[r] {
			l = r
			cur = nums[r]
		}
		if max < cur {
			out = nums[l : r+1]
			max = cur
		}
	}
	return out
}
