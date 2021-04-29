package main

import "math"

/*
  Given a non-negative integer c, decide whether there're two integers a and b such that a2 + b2 = c.

  Example 1:
    Input: c = 5
    Output: true
    Explanation: 1 * 1 + 2 * 2 = 5

  Example 2:
    Input: c = 3
    Output: false

  Example 3:
    Input: c = 4
    Output: true

  Example 4:
    Input: c = 2
    Output: true

  Example 5:
    Input: c = 1
    Output: true

  Constraints:
    0 <= c <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sum-of-square-numbers
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))
	for l <= r {
		sum := l*l + r*r
		if sum < c {
			l++
		} else if sum > c {
			r--
		} else {
			return true
		}
	}
	return false
}
