package main

/*
  给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
  （当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）

  示例 1:
    输入: N = 10
    输出: 9

  示例 2:
    输入: N = 1234
    输出: 1234

  示例 3:
    输入: N = 332
    输出: 299
    说明: N 是在 [0, 10^9] 范围内的一个整数。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/monotone-increasing-digits
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func monotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}
	nums := []int{}
	for n := N; n > 0; n /= 10 {
		nums = append(nums, n%10)
	}
	for i := 0; i < len(nums)>>1; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	index := 0
	for index < len(nums)-1 && nums[index] <= nums[index+1] {
		index++
	}
	if index == len(nums)-1 {
		return N
	}
	nums[index]--
	for j := index + 1; j < len(nums); j++ {
		nums[j] = 9
	}
	for k := index - 1; k >= 0 && nums[k] > nums[k+1]; k-- {
		nums[k]--
		nums[k+1] = 9
	}
	if nums[0] == 0 {
		nums = nums[1:]
	}
	out := 0
	for _, num := range nums {
		out = out*10 + num
	}
	return out
}
