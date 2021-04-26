package main

import "sort"

/*
  Given an integer array, return the k-th smallest distance among all the pairs. The distance of a pair (A, B) is defined as the absolute difference between A and B.

  Example 1:
    Input:
      nums = [1,3,1]
      k = 1
    Output: 0
    Explanation:
      Here are all the pairs:
      (1,3) -> 2
      (1,1) -> 0
      (3,1) -> 2
      Then the 1st smallest distance pair is (1,1), and its distance is 0.

  Note:
    1. 2 <= len(nums) <= 10000.
    2. 0 <= nums[i] < 1000000.
    3. 1 <= k <= len(nums) * (len(nums) - 1) / 2.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-k-th-smallest-pair-distance
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r := 0, int(1e6)
	for l < r {
		dist := (l + r) >> 1
		if countSmaller(nums, n, dist) < k {
			l = dist + 1
		} else {
			r = dist
		}
	}
	return l
}

func countSmaller(nums []int, n, dist int) int {
	out := 0
	for l, r := 0, 1; r < n; r++ {
		for nums[r]-nums[l] > dist {
			l++
		}
		out += r - l
	}
	return out
}
