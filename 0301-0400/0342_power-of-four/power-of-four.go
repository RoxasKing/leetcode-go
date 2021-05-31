package main

/*
  Given an integer n, return true if it is a power of four. Otherwise, return false.

  An integer n is a power of four, if there exists an integer x such that n == 4x.

  Example 1:
    Input: n = 16
    Output: true

  Example 2:
    Input: n = 5
    Output: false

  Example 3:
    Input: n = 1
    Output: true

  Constraints:
    -2^31 <= n <= 2^31 - 1

  Follow up: Could you solve it without loops/recursion?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/power-of-four
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Manipulation
// Math

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}
