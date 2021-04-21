package main

import "sort"

/*
  Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

  If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).

  The replacement must be in place and use only constant extra memory.

  Example 1:
    Input: nums = [1,2,3]
    Output: [1,3,2]

  Example 2:
    Input: nums = [3,2,1]
    Output: [1,2,3]

  Example 3:
    Input: nums = [1,1,5]
    Output: [1,5,1]

  Example 4:
    Input: nums = [1]
    Output: [1]

  Constraints:
    1. 1 <= nums.length <= 100
    2. 0 <= nums[i] <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/next-permutation
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}
	reverse(nums[i:])
	if i == 0 {
		return
	}
	j := sort.Search(n-i, func(j int) bool { return nums[j+i] > nums[i-1] }) + i
	nums[i-1], nums[j] = nums[j], nums[i-1]
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n>>1; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
