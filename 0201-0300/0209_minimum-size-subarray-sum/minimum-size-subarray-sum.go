package main

import "sort"

/*
  Given an array of positive integers nums and a positive integer target, return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than or equal to target. If there is no such subarray, return 0 instead.

  Example 1:
    Input: target = 7, nums = [2,3,1,2,4,3]
    Output: 2
    Explanation: The subarray [4,3] has the minimal length under the problem constraint.

  Example 2:
    Input: target = 4, nums = [1,4,4]
    Output: 1

  Example 3:
    Input: target = 11, nums = [1,1,1,1,1,1,1,1]
    Output: 0

  Constraints:
    1. 1 <= target <= 10^9
    2. 1 <= nums.length <= 10^5
    3. 1 <= nums[i] <= 10^5

  Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	out := n + 1
	for l, r, sum := 0, 0, 0; r < n; r++ {
		sum += nums[r]
		for l <= r && sum >= target {
			out = Min(out, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if out == n+1 {
		return 0
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Binary Seearch + Prefix Sum
func minSubArrayLen2(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	if sums[len(sums)-1] < s {
		return 0
	}
	min := 1<<31 - 1
	for l := range sums {
		target := s + sums[l]
		if target > sums[len(sums)-1] {
			break
		}
		r := sort.SearchInts(sums, target)
		min = Min(min, r-l)
	}
	if min == 1<<31-1 {
		return 0
	}
	return min
}
