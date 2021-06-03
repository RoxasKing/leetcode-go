package main

/*
  You are given an integer array nums. The absolute sum of a subarray [numsl, numsl+1, ..., numsr-1, numsr] is abs(numsl + numsl+1 + ... + numsr-1 + numsr).

  Return the maximum absolute sum of any (possibly empty) subarray of nums.

  Note that abs(x) is defined as follows:

    1. If x is a negative integer, then abs(x) = -x.
    2. If x is a non-negative integer, then abs(x) = x.

  Example 1:
    Input: nums = [1,-3,2,3,-4]
    Output: 5
    Explanation: The subarray [2,3] has absolute sum = abs(2+3) = abs(5) = 5.

  Example 2:
    Input: nums = [2,-5,1,-4,3,-2]
    Output: 8
    Explanation: The subarray [-5,1,-4] has absolute sum = abs(-5+1-4) = abs(-8) = 8.

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. -10^4 <= nums[i] <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-absolute-sum-of-any-subarray
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy Algoritm

func maxAbsoluteSum(nums []int) int {
	out := 0
	max, min := -1<<31, 1<<31-1
	for _, num := range nums {
		max = Max(max+num, num)
		min = Min(min+num, num)
		out = Max(out, Max(Abs(max), Abs(min)))
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
