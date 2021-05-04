package main

/*
  Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product, and return the product.

  It is guaranteed that the answer will fit in a 32-bit integer.

  A subarray is a contiguous subsequence of the array.

  Example 1:
    Input: nums = [2,3,-2,4]
    Output: 6
    Explanation: [2,3] has the largest product 6.

  Example 2:
    Input: nums = [-2,0,-1]
    Output: 0
    Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

  Constraints:
    1. 1 <= nums.length <= 2 * 10^4
    2. -10 <= nums[i] <= 10
    3. The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-product-subarray
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming

func maxProduct(nums []int) int {
	max, min, out := nums[0], nums[0], nums[0]
	for _, num := range nums[1:] {
		a, b := max*num, min*num
		max = Max(num, Max(a, b))
		min = Min(num, Min(a, b))
		out = Max(out, max)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
