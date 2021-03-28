package main

/*
  You are given three positive integers n, index and maxSum. You want to construct an array nums (0-indexed) that satisfies the following conditions:
    1. nums.length == n
    2. nums[i] is a positive integer where 0 <= i < n.
    3. abs(nums[i] - nums[i+1]) <= 1 where 0 <= i < n-1.
    4. The sum of all the elements of nums does not exceed maxSum.
    5.nums[index] is maximized.
  Return nums[index] of the constructed array.

  Note that abs(x) equals x if x >= 0, and -x otherwise.

  Example 1:
    Input: n = 4, index = 2,  maxSum = 6
    Output: 2
    Explanation: The arrays [1,1,2,1] and [1,2,2,1] satisfy all the conditions. There are no other valid arrays with a larger value at the given index.

  Example 2:
    Input: n = 6, index = 1,  maxSum = 10
    Output: 3

  Constraints:
    1. 1 <= n <= maxSum <= 10^9
    2. 0 <= index < n

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-value-at-a-given-index-in-a-bounded-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy Algorithm + Binary Search + Math
func maxValue(n int, index int, maxSum int) int {
	maxSum -= n
	lLen := index
	rLen := n - (index + 1)
	l, r := 0, maxSum
	for l < r {
		m := (l + r) / 2
		lSum := Min(lLen, m) * (m - 1 + m - 1 - (Min(lLen, m) - 1)) / 2 // 1+2+...+Min(lLen, m-1)
		rSum := Min(rLen, m) * (m - 1 + m - 1 - (Min(rLen, m) - 1)) / 2 // 1+2+...+Min(rLen, m-1)
		if lSum+m+rSum < maxSum {
			l = m + 1
		} else {
			r = m
		}
	}

	lSum := Min(lLen, l) * (l - 1 + l - 1 - (Min(lLen, l) - 1)) / 2
	rSum := Min(rLen, l) * (l - 1 + l - 1 - (Min(rLen, l) - 1)) / 2
	if maxSum <= lSum+rSum+l-1 {
		return l
	}
	return l + 1
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
