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
	out := 0
	n := len(nums)
	for l, r := 0, 0; r < n; r++ {
		target -= nums[r]
		for l < r && target+nums[l] <= 0 {
			target += nums[l]
			l++
		}
		if target <= 0 && (out == 0 || r+1-l < out) {
			out = r + 1 - l
		}
	}
	return out
}

// Binary Seearch + Prefix Sum
func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + nums[i]
	}

	out := 0
	for i := 0; i < n; i++ {
		j := sort.Search((n+1)-(i+1), func(j int) bool { return pSum[j+(i+1)] >= pSum[i]+target }) + (i + 1)
		if j > n {
			break
		}
		if out == 0 || j-i < out {
			out = j - i
		}
	}
	return out
}
