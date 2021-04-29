package main

/*
  Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

  Example 1:
    Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
    Output: 6
    Explanation: [4,-1,2,1] has the largest sum = 6.

  Example 2:
    Input: nums = [1]
    Output: 1

  Example 3:
    Input: nums = [5,4,-1,7,8]
    Output: 23

  Constraints:
    1. 1 <= nums.length <= 3 * 10^4
    2. -10^5 <= nums[i] <= 10^5

  Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-subarray
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func maxSubArray(nums []int) int {
	out := -1 << 31
	cur := -1 << 31
	for _, num := range nums {
		cur = Max(cur+num, num)
		out = Max(out, cur)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
