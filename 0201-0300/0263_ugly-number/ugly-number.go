package main

/*
  Given an integer n, return true if n is an ugly number.

  Ugly number is a positive number whose prime factors only include 2, 3, and/or 5.

  Example 1:
    Input: n = 6
    Output: true
    Explanation: 6 = 2 × 3

  Example 2:
    Input: n = 8
    Output: true
    Explanation: 8 = 2 × 2 × 2

  Example 3:
    Input: n = 14
    Output: false
    Explanation: 14 is not ugly since it includes another prime factor 7.

  Example 4:
    Input: n = 1
    Output: true
    Explanation: 1 is typically treated as an ugly number.

  Constraints:
    -2^31 <= n <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ugly-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}
