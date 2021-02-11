package main

/*
  Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most one element.
  We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).

  Example 1:
    Input: nums = [4,2,3]
    Output: true
    Explanation: You could modify the first 4 to 1 to get a non-decreasing array.

  Example 2:
    Input: nums = [4,2,1]
    Output: false
    Explanation: You can't get a non-decreasing array by modify at most one element.

  Constraints:
    n == nums.length
    1 <= n <= 10^4
    -10^5 <= nums[i] <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/non-decreasing-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func checkPossibility(nums []int) bool {
	n := len(nums)
	changed := false
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			if changed {
				return false
			}
			if i-1 >= 0 {
				if nums[i+1] >= nums[i-1] {
					nums[i] = nums[i-1]
				} else {
					nums[i+1] = nums[i]
				}
			} else {
				nums[i] = nums[i+1]
			}
			changed = true
		}
	}
	return true
}
