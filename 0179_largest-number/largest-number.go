package main

import (
	"sort"
	"strconv"
	"strings"
)

/*
  Given a list of non-negative integers nums, arrange them such that they form the largest number.

  Note: The result may be very large, so you need to return a string instead of an integer.

  Example 1:
    Input: nums = [10,2]
    Output: "210"

  Example 2:
    Input: nums = [3,30,34,5,9]
    Output: "9534330"

  Example 3:
    Input: nums = [1]
    Output: "1"

  Example 4:
    Input: nums = [10]
    Output: "10"

  Constraints:
    1. 1 <= nums.length <= 100
    2. 0 <= nums[i] <= 109

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/largest-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	if strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}
