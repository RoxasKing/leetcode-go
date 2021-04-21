package main

/*
  You are given an array nums that consists of non-negative integers. Let us define rev(x) as the reverse of the non-negative integer x. For example, rev(123) = 321, and rev(120) = 21. A pair of indices (i, j) is nice if it satisfies all of the following conditions:
    1. 0 <= i < j < nums.length
    2. nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
  Return the number of nice pairs of indices. Since that number can be too large, return it modulo 109 + 7.

  Example 1:
    Input: nums = [42,11,1,97]
    Output: 2
    Explanation: The two pairs are:
     - (0,3) : 42 + rev(97) = 42 + 79 = 121, 97 + rev(42) = 97 + 24 = 121.
     - (1,2) : 11 + rev(1) = 11 + 1 = 12, 1 + rev(11) = 1 + 11 = 12.

  Example 2:
    Input: nums = [13,10,35,24,76]
    Output: 4

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. 0 <= nums[i] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-nice-pairs-in-an-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func countNicePairs(nums []int) int {
	cnt := make(map[int]int)
	for i := range nums {
		cnt[nums[i]-rev(nums[i])]++
	}
	out := 0
	for _, x := range cnt {
		res := (x * (x - 1) / 2) % (1e9 + 7)
		out = (out + res) % (1e9 + 7)
	}
	return out
}

func rev(num int) int {
	out := 0
	for num > 0 {
		out = out*10 + num%10
		num /= 10
	}
	return out
}
