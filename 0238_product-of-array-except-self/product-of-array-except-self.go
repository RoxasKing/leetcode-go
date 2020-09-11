package main

/*
  给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

  提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。

  说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

  进阶：
  你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/product-of-array-except-self
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	form := make([]int, n)
	reve := make([]int, n)
	form[0] = nums[0]
	reve[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		form[i] = form[i-1] * nums[i]
	}
	for i := n - 2; i >= 0; i-- {
		reve[i] = reve[i+1] * nums[i]
	}
	out := make([]int, n)
	out[0] = reve[1]
	out[n-1] = form[n-2]
	for i := 1; i < n-1; i++ {
		out[i] = form[i-1] * reve[i+1]
	}
	return out
}

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	out[0] = 1
	for i := 1; i < n; i++ {
		out[i] = out[i-1] * nums[i-1]
	}
	R := 1
	for i := n - 1; i >= 0; i-- {
		out[i] *= R
		R *= nums[i]
	}
	return out
}
