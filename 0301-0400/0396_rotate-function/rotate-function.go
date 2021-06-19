package main

/*
  You are given an integer array nums of length n.

  Assume arrk to be an array obtained by rotating nums by k positions clock-wise. We define the rotation function F on nums as follow:

    F(k) = 0 * arrk[0] + 1 * arrk[1] + ... + (n - 1) * arrk[n - 1].

  Return the maximum value of F(0), F(1), ..., F(n-1).

  Example 1:
    Input: nums = [4,3,2,6]
    Output: 26
    Explanation:
      F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
      F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
      F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
      F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26
      So the maximum value of F(0), F(1), F(2), F(3) is F(3) = 26.

  Example 2:
    Input: nums = [1000000007]
    Output: 0

  Constraints:
    1. n == nums.length
    2. 1 <= n <= 3 * 10^4
    3. -2^31 <= nums[i] <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/rotate-function
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math

func maxRotateFunction(nums []int) int {
	n := len(nums)
	sum := 0
	f := 0
	for i, num := range nums {
		sum += num
		f += i * num
	}

	out := f
	for i := 1; i < n; i++ {
		f -= sum - nums[i-1]
		f += nums[i-1] * (n - 1)
		out = Max(out, f)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
