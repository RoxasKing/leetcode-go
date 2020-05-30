package codinginterviews

/*
  输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
  要求时间复杂度为O(n)。

  提示：
    1 <= arr.length <= 10^5
    -100 <= arr[i] <= 100
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := -1 << 31
	cur := -1 << 31
	for _, num := range nums {
		cur = Max(cur+num, num)
		max = Max(max, cur)
	}
	return max
}
