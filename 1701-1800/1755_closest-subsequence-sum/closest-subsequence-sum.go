package main

import (
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

// TODO

// Important!

// Divide-and-Conquer Algorithm

func minAbsDifference(nums []int, goal int) int {
	out := Min(Abs(goal), Abs(nums[0]-goal))
	n := len(nums)
	if n == 1 {
		return out
	}

	m := n >> 1
	lp := partition(nums[:m])
	rp := partition(nums[m:])

	set := map[int]struct{}{}
	for num := range lp {
		out = Min(out, Abs(num-goal))
		set[num] = struct{}{}
	}
	arr := make([]int, 0, len(set))
	for num := range set {
		arr = append(arr, num)
	}
	sort.Ints(arr)

	for num := range rp {
		out = Min(out, Abs(num-goal))
		i := sort.Search(len(arr), func(i int) bool { return num+arr[i] >= goal })
		if i == len(arr) {
			out = Min(out, Abs(arr[i-1]+num-goal))
			continue
		}
		out = Min(out, Abs(arr[i]+num-goal))
		if i > 0 {
			out = Min(out, Abs(arr[i-1]+num-goal))
		}
	}
	return out
}

func partition(nums []int) map[int]struct{} {
	n := len(nums)
	if n == 1 {
		return map[int]struct{}{nums[0]: {}}
	}
	m := n >> 1
	lp := partition(nums[:m])
	rp := partition(nums[m:])
	set := map[int]struct{}{}
	for num1 := range lp {
		set[num1] = struct{}{}
		for num2 := range rp {
			set[num2] = struct{}{}
			set[num1+num2] = struct{}{}
		}
	}
	return set
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
