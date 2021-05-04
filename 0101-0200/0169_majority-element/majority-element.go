package main

/*
  Given an array nums of size n, return the majority element.

  The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

  Example 1:
    Input: nums = [3,2,3]
    Output: 3

  Example 2:
    Input: nums = [2,2,1,1,1,2,2]
    Output: 2

  Constraints:
    1. n == nums.length
    2. 1 <= n <= 5 * 10^4
    3. -2^31 <= nums[i] <= 2^31 - 1

  Follow-up: Could you solve the problem in linear time and in O(1) space?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/majority-element
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Boyer-Moore
func majorityElement(nums []int) int {
	cur, cnt := nums[0], 1
	for _, num := range nums[1:] {
		if num == cur {
			cnt++
		} else {
			cnt--
		}

		if cnt == 0 {
			cur, cnt = num, 1
		}
	}
	return cur
}
