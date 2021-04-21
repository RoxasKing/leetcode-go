package main

/*
  Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

  Follow up: Could you implement a solution with a linear runtime complexity and without using extra memory?

  Example 1:
    Input: nums = [2,2,1]
    Output: 1

  Example 2:
    Input: nums = [4,1,2,1,2]
    Output: 4

  Example 3:
    Input: nums = [1]
    Output: 1

  Constraints:
    1 <= nums.length <= 3 * 10^4
    -3 * 10^4 <= nums[i] <= 3 * 10^4
    Each element in the array appears twice except for one element which appears only once.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/single-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Manipulation
func singleNumber(nums []int) int {
	out := 0
	for _, num := range nums {
		out ^= num
	}
	return out
}
