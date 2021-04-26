package main

import "sort"

/*
  The frequency of an element is the number of times it occurs in an array.

  You are given an integer array nums and an integer k. In one operation, you can choose an index of nums and increment the element at that index by 1.

  Return the maximum possible frequency of an element after performing at most k operations.

  Example 1:
    Input: nums = [1,2,4], k = 5
    Output: 3
    Explanation: Increment the first element three times and the second element two times to make nums = [4,4,4].
      4 has a frequency of 3.

  Example 2:
    Input: nums = [1,4,8,13], k = 5
    Output: 2
    Explanation: There are multiple optimal solutions:
      - Increment the first element three times to make nums = [4,4,8,13]. 4 has a frequency of 2.
      - Increment the second element four times to make nums = [1,8,8,13]. 8 has a frequency of 2.
      - Increment the third element five times to make nums = [1,4,13,13]. 13 has a frequency of 2.

  Example 3:
    Input: nums = [3,9,6], k = 2
    Output: 1

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. 1 <= nums[i] <= 10^5
    3. 1 <= k <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/frequency-of-the-most-frequent-element
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func maxFrequency(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	n := len(nums)
	out := 0

	for l, r := 0, 0; r < n; r++ {
		k -= nums[l] - nums[r]
		for l < r && k < 0 {
			k += (r - l) * (nums[l] - nums[l+1])
			l++
		}
		out = Max(out, r+1-l)
	}

	return out
}

// Binary Search
func maxFrequency2(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	n := len(nums)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + nums[i]
	}
	out := 0
	for i := 1; i <= n; i++ {
		j := sort.Search(n+1-(i+1), func(j int) bool {
			j += i + 1
			return (pSum[j]-pSum[i-1]+k)/(j-i+1) < nums[i-1]
		}) + (i + 1)
		out = Max(out, j-i)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
