package main

/*
  Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.

  Example 1:
    Input: nums = [0,1]
    Output: 2
    Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.

  Example 2:
    Input: nums = [0,1,0]
    Output: 2
    Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. nums[i] is either 0 or 1.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/contiguous-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Hash

func findMaxLength(nums []int) int {
	out := 0
	dict := map[int]int{0: -1}
	diff := 0
	for i, num := range nums {
		if num == 0 {
			diff--
		} else {
			diff++
		}
		if j, ok := dict[diff]; ok {
			out = Max(out, i-j)
		} else {
			dict[diff] = i
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
