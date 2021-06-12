package main

/*
  Given integer array nums, return the third maximum number in this array. If the third maximum does not exist, return the maximum number.

  Example 1:
    Input: nums = [3,2,1]
    Output: 1
    Explanation: The third maximum is 1.

  Example 2:
    Input: nums = [1,2]
    Output: 2
    Explanation: The third maximum does not exist, so the maximum (2) is returned instead.

  Example 3:
    Input: nums = [2,2,3,1]
    Output: 1
    Explanation: Note that the third maximum here means the third maximum distinct number.
      Both numbers with value 2 are both considered as second maximum.

  Constraints:
    1. 1 <= nums.length <= 10^4
    2. -2^31 <= nums[i] <= 2^31 - 1

    Follow up: Can you find an O(n) solution?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/third-maximum-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func thirdMax(nums []int) int {
	a, b, c := -1<<63, -1<<63, -1<<63
	for _, num := range nums {
		if num > a {
			a, b, c = num, a, b
		} else if num < a && num > b {
			b, c = num, b
		} else if num < b && num > c {
			c = num
		}
	}
	if c == -1<<63 {
		return a
	}
	return c
}
