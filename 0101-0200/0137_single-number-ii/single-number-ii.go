package main

/*
  Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.

  Example 1:
    Input: nums = [2,2,3,2]
    Output: 3

  Example 2:
    Input: nums = [0,1,0,1,0,1,99]
    Output: 99

  Constraints:
    1. 1 <= nums.length <= 3 * 10^4
    2. -2^31 <= nums[i] <= 2^31 - 1
    3. Each element in nums appears exactly three times except for one element which appears once.

  Follow up: Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/single-number-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Manipulation
func singleNumber(nums []int) int {
	var s1, s2 int
	for _, num := range nums {
		s1 = ^s2 & (s1 ^ num)
		s2 = ^s1 & (s2 ^ num)
	}
	return s1
}
