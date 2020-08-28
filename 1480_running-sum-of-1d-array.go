package main

/*
  给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
  请返回 nums 的动态和。

  提示：
    1 <= nums.length <= 1000
    -10^6 <= nums[i] <= 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/running-sum-of-1d-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func runningSum(nums []int) []int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	return sums
}
