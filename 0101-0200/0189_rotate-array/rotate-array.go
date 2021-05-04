package main

/*
Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:
  Input: nums = [1,2,3,4,5,6,7], k = 3
  Output: [5,6,7,1,2,3,4]
  Explanation:
    rotate 1 steps to the right: [7,1,2,3,4,5,6]
    rotate 2 steps to the right: [6,7,1,2,3,4,5]
    rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
  Input: nums = [-1,-100,3,99], k = 2
  Output: [3,99,-1,-100]
  Explanation:
    rotate 1 steps to the right: [99,-1,-100,3]
    rotate 2 steps to the right: [3,99,-1,-100]

Constraints:
  1. 1 <= nums.length <= 10^5
  2. -2^31 <= nums[i] <= 2^31 - 1
  3. 0 <= k <= 105

Follow up:
  1. Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
  2. Could you do it in-place with O(1) extra space?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n>>1; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
