package main

/*
  You are given an integer array nums. The unique elements of an array are the elements that appear exactly once in the array.

  Return the sum of all the unique elements of nums.

  Example 1:
    Input: nums = [1,2,3,2]
    Output: 4
    Explanation: The unique elements are [1,3], and the sum is 4.

  Example 2:
    Input: nums = [1,1,1,1,1]
    Output: 0
    Explanation: There are no unique elements, and the sum is 0.

  Example 3:
    Input: nums = [1,2,3,4,5]
    Output: 15
    Explanation: The unique elements are [1,2,3,4,5], and the sum is 15.

  Constraints:
    1. 1 <= nums.length <= 100
    2. 1 <= nums[i] <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sum-of-unique-elements
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func sumOfUnique(nums []int) int {
	out := 0
	cnt := [101]int{}
	for _, num := range nums {
		if cnt[num] == 0 {
			out += num
		} else if cnt[num] == 1 {
			out -= num
		}
		cnt[num]++
	}
	return out
}
