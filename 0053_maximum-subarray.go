package leetcode

/*
  给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/

func maxSubArray(nums []int) int {
	sum, maxSum := -1<<31, -1<<31
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	for _, num := range nums {
		sum = max(sum+num, num)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func maxSubArrayReturnArray(nums []int) []int {
	sum, maxSum := -1<<31, -1<<31
	var left, right, l, r int
	for right = range nums {
		tmp := sum + nums[right]
		if tmp < nums[right] {
			left = right
			sum = nums[right]
		} else {
			sum = tmp
		}
		if maxSum < sum {
			l, r = left, right+1
			maxSum = sum
		}
	}
	return nums[l:r]
}
