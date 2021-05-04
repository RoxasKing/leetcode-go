package main

/*
  Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

  Note that you must do this in-place without making a copy of the array.

  Example 1:
    Input: nums = [0,1,0,3,12]
    Output: [1,3,12,0,0]

  Example 2:
    Input: nums = [0]
    Output: [0]

  Constraints:
    1. 1 <= nums.length <= 10^4
    2. -2^31 <= nums[i] <= 2^31 - 1

  Follow up: Could you minimize the total number of operations done?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/move-zeroes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func moveZeroes(nums []int) {
	index := 0
	for _, num := range nums {
		if num != 0 {
			nums[index] = num
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}
