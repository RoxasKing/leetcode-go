package main

/*
  Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

  We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

  Example 1:
    Input: nums = [2,0,2,1,1,0]
    Output: [0,0,1,1,2,2]

  Example 2:
    Input: nums = [2,0,1]
    Output: [0,1,2]

  Example 3:
    Input: nums = [0]
    Output: [0]

  Example 4:
    Input: nums = [1]
    Output: [1]

  Constraints:
    1. n == nums.length
    2. 1 <= n <= 300
    3. nums[i] is 0, 1, or 2.

  Follow up:
    1. Could you solve this problem without using the library's sort function?
    2. Could you come up with a one-pass algorithm using only O(1) constant space?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sort-colors
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Two Pointers

func sortColors(nums []int) {
	n := len(nums)
	l, m, r := 0, 0, n-1
	for m <= r {
		if nums[m] == 0 {
			nums[l], nums[m] = nums[m], nums[l]
			l++
			m++
		} else if nums[m] == 2 {
			nums[m], nums[r] = nums[r], nums[m]
			r--
		} else {
			m++
		}
	}
}
