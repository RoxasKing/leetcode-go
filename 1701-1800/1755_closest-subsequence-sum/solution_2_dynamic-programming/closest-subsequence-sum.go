package main

import (
	"fmt"
	"sort"
)

/*
  You are given an integer array nums and an integer goal.

  You want to choose a subsequence of nums such that the sum of its elements is the closest possible to goal. That is, if the sum of the subsequence's elements is sum, then you want to minimize the absolute difference abs(sum - goal).

  Return the minimum possible value of abs(sum - goal).

  Note that a subsequence of an array is an array formed by removing some elements (possibly all or none) of the original array.

  Example 1:
    Input: nums = [5,-7,3,5], goal = 6
    Output: 0
    Explanation: Choose the whole array as a subsequence, with a sum of 6.
      This is equal to the goal, so the absolute difference is 0.

  Example 2:
    Input: nums = [7,-9,15,-2], goal = -5
    Output: 1
    Explanation: Choose the subsequence [7,-9,-2], with a sum of -4.
      The absolute difference is abs(-4 - (-5)) = abs(1) = 1, which is the minimum.

  Example 3:
    Input: nums = [1,2,3], goal = -7
    Output: 7

  Constraints:
    1. 1 <= nums.length <= 40
    2. -10^7 <= nums[i] <= 10^7
    3. -10^9 <= goal <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/closest-subsequence-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming
// Two Pointers

func minAbsDifference(nums []int, goal int) int {
	out := Min(Abs(goal), Abs(nums[0]-goal))
	n := len(nums)
	if n == 1 {
		return out
	}

	m := n >> 1
	ls, rs := m, n-m
	lSum := make([]int, 1<<ls)
	rSum := make([]int, 1<<rs)
	for i := 1; i < 1<<ls; i++ {
		for j := 0; j < ls; j++ {
			if i&(1<<j) > 0 {
				lSum[i] = lSum[i-(1<<j)] + nums[j]
				break
			}
		}
	}
	for i := 1; i < 1<<rs; i++ {
		for j := 0; j < rs; j++ {
			if i&(1<<j) > 0 {
				rSum[i] = rSum[i-(1<<j)] + nums[ls+j]
				break
			}
		}
	}

	for _, sum := range lSum {
		out = Min(out, Abs(sum-goal))
	}
	for _, sum := range rSum {
		out = Min(out, Abs(sum-goal))
	}
	sort.Ints(lSum)
	sort.Ints(rSum)

	fmt.Println(lSum)
	fmt.Println(rSum)

	for i, j := 0, 1<<rs-1; i < 1<<ls && j >= 0; {
		sum := lSum[i] + rSum[j]
		out = Min(out, Abs(sum-goal))
		if sum <= goal {
			i++
		} else {
			j--
		}
	}
	return out
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
