package main

/*
  You are given an array of integers nums (0-indexed) and an integer k.

  The score of a subarray (i, j) is defined as min(nums[i], nums[i+1], ..., nums[j]) * (j - i + 1). A good subarray is a subarray where i <= k <= j.

  Return the maximum possible score of a good subarray.

  Example 1:
    Input: nums = [1,4,3,7,4,5], k = 3
    Output: 15
    Explanation: The optimal subarray is (1, 5) with a score of min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15.

  Example 2:
    Input: nums = [5,5,4,5,4,1,1,1], k = 0
    Output: 20
    Explanation: The optimal subarray is (0, 4) with a score of min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20.

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. 1 <= nums[i] <= 2 * 10^4
    3. 0 <= k < nums.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-score-of-a-good-subarray
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func maximumScore(nums []int, k int) int {
	n := len(nums)
	out := 0
	for min, l, r := nums[k], k, k; min >= 1; min-- {
		for l-1 >= 0 && nums[l-1] >= min {
			l--
		}
		for r+1 <= n-1 && nums[r+1] >= min {
			r++
		}
		out = Max(out, min*(r-l+1))
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
